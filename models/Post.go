package models

import (
	"bytes"
	"github.com/jackc/pgx"
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
	"strconv"
	"time"
	"context"
	"github.com/jackc/pgx/pgtype"
)

//easyjson:json
type Post struct {
	ID       int        `json:"id"`
	Author   string     `json:"author"`
	Created  *time.Time `json:"created"`
	ForumID  string     `json:"forum"`
	IsEdited bool       `json:"isEdited"`
	Message  string     `json:"message"`
	Parent   int        `json:"parent,omitempty"`
	Thread   int        `json:"thread"`
	Slug     string     `json:"slug,omitempty"`
	ParentId []int64
}

//easyjson:json
type PostDetail struct {
	Author *User   `json:"author,omitempty"`
	Forum  *Forum  `json:"forum,omitempty"`
	Post   *Post   `json:"post,omitempty"`
	Thread *Thread `json:"thread,omitempty"`
}

//easyjson:json
type PostUpdate struct {
	Message string `json:"message"`
}

//easyjson:json
type Posts []*Post

func (posts *Posts) PostsCreate(slug string) error {
	tx := config.StartTransaction()
	defer tx.Rollback()

	batch := tx.BeginBatch()
	defer batch.Close()

	//checking thread id or slug
	var forumSlug string

	id, err := strconv.Atoi(slug)
	if err != nil {
		//	id is slug (string)
		if err = tx.QueryRow(helpers.SelectThreadIdForumSlug, slug).Scan(&id, &forumSlug); err != nil {
			return errors.ThreadNotFound
		}
	} else {
		if err = tx.QueryRow(helpers.SelectThreadIdForumSlugByID, id).Scan(&id, &forumSlug); err != nil {
			return errors.ThreadNotFound
		}
	}

	if len(*posts) == 0 {
		return errors.NoPostsForCreate
	}

	created, _ := time.Parse("2006-01-02T15:04:05.000000Z", "2006-01-02T15:04:05.010000Z")

	//NEW
	users := Users{}
	//authors := make(map[string]struct{})
	var nonRootPosts []int

	for i, post := range *posts {
		if post.Parent != 0 {
			nonRootPosts = append(nonRootPosts, i)
			batch.Queue(helpers.SelectThreadID, []interface{}{post.Parent, id},
				[]pgtype.OID{pgtype.Int4OID, pgtype.Int4OID}, nil)
		}
	}

	for _, post := range *posts {
		batch.Queue(helpers.SelectUserNick, []interface{}{post.Author},
			[]pgtype.OID{pgtype.TextOID}, nil)
	}

	if err = batch.Send(context.Background(), nil); err != nil {
		log.Panic(err)
	}

	for _, val := range nonRootPosts {
		var parents []int64
		if err := batch.QueryRowResults().Scan(&parents); err != nil {
			return errors.NoThreadParent
		}
		(*posts)[val].ParentId = append((*posts)[val].ParentId, parents...)
	}

	//check authors
	for range *posts {
		user := User{}
		if err := batch.QueryRowResults().Scan(&user.About,
			&user.Email, &user.FullName, &user.NickName); err != nil {
				return errors.ThreadNotFound
		}
		users = append(users, &user)
	}

	//TODO: сделать через batch
	//for i, post := range *posts {
	//	log.Println("CREATE ", i)
	//	batch.Queue(helpers.CreatePost, []interface{}{&users[i].NickName, &created,
	//		&forumSlug, &post.Message, &post.Parent, &id}, []pgtype.OID{
	//		pgtype.TextOID, pgtype.TimestamptzOID, pgtype.TextOID, pgtype.TextOID,
	//		pgtype.Int4OID, pgtype.Int4OID}, nil)
	//}

	//for _, user := range users {
	//	batch.Queue(helpers.InsertForumUsers, []interface{}{&user.About, &user.Email,
	//		&user.FullName, &user.NickName, &forumSlug}, []pgtype.OID{
	//		pgtype.TextOID, pgtype.TextOID, pgtype.TextOID, pgtype.TextOID, pgtype.TextOID}, nil)
	//}

	//if err = batch.Send(context.Background(), nil); err != nil {
	//	log.Panic(err)
	//}

	//for _, post := range *posts {
	//	if err := batch.QueryRowResults().Scan(&post.ID); err != nil {
	//		log.Println(err)
	//		return errors.ThreadNotFound
	//	}
	//}

	//log.Println(len(users))
	//for range users {
	//	if _, err := batch.QueryResults(); err != nil {
	//		log.Panic(err)
	//	}
	//}
	//OLD

	for i, post := range *posts {

		if err = tx.QueryRow(helpers.CreatePost, &users[i].NickName, &created,
			&forumSlug, &post.Message, &post.Parent, &id).Scan(&post.ID); err != nil {
			return errors.ThreadNotFound
		}

		post.ParentId = append(post.ParentId, int64(post.ID))

		_, err = tx.Exec(helpers.InsertForumUsers, &users[i].About, &users[i].Email,
			&users[i].FullName, &users[i].NickName, &forumSlug)

		_, err = tx.Exec(helpers.CreatePostParent, &post.ID, &post.ParentId)
		if err != nil {
			log.Println(err)
		}
		post.Created = &created
		post.IsEdited = false
		post.ForumID = forumSlug
		post.Thread = id

	}

	tx.Exec(helpers.InsertForumPostCnt, len(*posts), &forumSlug)

	tx.Commit()
	return nil
}

func GetPostThreadId(slug string) (int, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	//checking thread id or slug
	var forumSlug string

	id, err := strconv.Atoi(slug)
	if err != nil {
		//	id is slug (string)
		if err = tx.QueryRow(helpers.SelectThreadIdForumSlug, slug).Scan(&id, &forumSlug); err != nil {
			return id, errors.ThreadNotFound
		}
	} else {
		if err = tx.QueryRow(helpers.SelectThreadIdForumSlugByID, id).Scan(&id, &forumSlug); err != nil {
			return id, errors.ThreadNotFound
		}
	}

	return id, nil
}

func GetPostsSortFlat(threadId int, limit []byte,
	since []byte, desc []byte) (*Posts, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	posts := Posts{}
	var err error
	var result *pgx.Rows
	if since != nil {
		if bytes.Equal([]byte("true"), desc) {
			result, err = tx.Query(helpers.SelectPostsSinceFlatDesc, &threadId,
				string(since), string(limit))
		} else {
			result, err = tx.Query(helpers.SelectPostsSinceFlat, &threadId,
				string(since), string(limit))
		}
	} else {
		if bytes.Equal([]byte("true"), desc) {
			result, err = tx.Query(helpers.SelectPostsFlatDesc, &threadId, string(limit))
		} else {
			result, err = tx.Query(helpers.SelectPostsFlat, &threadId, string(limit))
		}
	}
	defer result.Close()

	if err != nil {
		return nil, errors.ThreadNotFound
	}

	for result.Next() {
		post := Post{}

		err = result.Scan(&post.Author, &post.Created, &post.ForumID,
			&post.ID, &post.IsEdited, &post.Message, &post.Parent, &post.Thread)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &post)
	}

	if len(posts) == 0 {
		var cnt int
		if err = tx.QueryRow("SELECT 1 FROM thread WHERE id = $1", &threadId).Scan(&cnt); err != nil {
			return nil, errors.ThreadNotFound
		}
	}

	return &posts, nil
}

func GetPostsSortTree(threadId int, limit []byte,
	since []byte, desc []byte) (*Posts, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	posts := Posts{}
	var err error
	var result *pgx.Rows
	if since != nil {
		if bytes.Equal([]byte("true"), desc) {
			result, err = tx.Query(helpers.SelectPostsSinceTreeDesc, &threadId,
				string(since), string(limit))
		} else {
			result, err = tx.Query(helpers.SelectPostsSinceTree, &threadId,
				string(since), string(limit))
		}
	} else {
		if bytes.Equal([]byte("true"), desc) {
			result, err = tx.Query(helpers.SelectPostsTreeDesc, &threadId, string(limit))
		} else {
			result, err = tx.Query(helpers.SelectPostsTree, &threadId, string(limit))
		}
	}
	defer result.Close()

	if err != nil {
		return nil, errors.ThreadNotFound
	}

	for result.Next() {
		post := Post{}

		err = result.Scan(&post.Author, &post.Created, &post.ForumID,
			&post.ID, &post.IsEdited, &post.Message, &post.Parent, &post.Thread)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &post)
	}

	if len(posts) == 0 {
		var cnt int
		if err = tx.QueryRow("SELECT 1 FROM thread WHERE id = $1", &threadId).Scan(&cnt); err != nil {
			return nil, errors.ThreadNotFound
		}
	}

	return &posts, nil
}

func GetPostsSortParentTree(threadId int, limit []byte,
	since []byte, desc []byte) (*Posts, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	posts := Posts{}
	var err error
	var result *pgx.Rows
	if since != nil {
		if limit != nil {
			if bytes.Equal([]byte("true"), desc) {
				result, err = tx.Query(helpers.SelectPostsSinceParentTreeLimitDesc, &threadId,
					string(limit), string(since))
			} else {
				result, err = tx.Query(helpers.SelectPostsSinceParentTreeLimit, &threadId,
					string(limit), string(since))
			}
		} else {
			if bytes.Equal([]byte("true"), desc) {
				result, err = tx.Query(helpers.SelectPostsSinceParentTreeDesc, &threadId,
					string(since))
			} else {
				result, err = tx.Query(helpers.SelectPostsSinceParentTree, &threadId,
					string(since))
			}
		}
	} else {
		if limit != nil {
			if bytes.Equal([]byte("true"), desc) {
				result, err = tx.Query(helpers.SelectPostsParentTreeLimitDesc, &threadId, string(limit))
			} else {
				result, err = tx.Query(helpers.SelectPostsParentTreeLimit, &threadId, string(limit))
			}
		} else {
			if bytes.Equal([]byte("true"), desc) {
				result, err = tx.Query(helpers.SelectPostsParentTreeDesc, &threadId)
			} else {
				result, err = tx.Query(helpers.SelectPostsParentTree, &threadId)
			}
		}
	}
	defer result.Close()

	if err != nil {
		return nil, errors.ThreadNotFound
	}

	for result.Next() {
		post := Post{}

		err = result.Scan(&post.Author, &post.Created, &post.ForumID,
			&post.ID, &post.IsEdited, &post.Message, &post.Parent, &post.Thread)
		if err != nil {
			log.Fatal(err)
		}
		posts = append(posts, &post)
	}

	if len(posts) == 0 {
		var cnt int
		if err = tx.QueryRow("SELECT 1 FROM thread WHERE id = $1", &threadId).Scan(&cnt); err != nil {
			return nil, errors.ThreadNotFound
		}
	}

	return &posts, nil
}

func PostDetails(id string, related []string) (*PostDetail, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	postDetail := PostDetail{}
	postDetail.Post = &Post{}

	postId, _ := strconv.Atoi(id)

	err := tx.QueryRow(helpers.SelectPost, postId).Scan(&postDetail.Post.Author,
		&postDetail.Post.Created, &postDetail.Post.ForumID, &postDetail.Post.Message,
		&postDetail.Post.Parent, &postDetail.Post.Thread, &postDetail.Post.IsEdited)
	if err != nil {
		return nil, errors.PostNotFound
	}
	postDetail.Post.ID = postId

	if related == nil {
		return &postDetail, nil
	}
	for _, val := range related {
		switch val {
		case "user":
			postDetail.Author = &User{}
			tx.QueryRow(helpers.SelectUserProfile, &postDetail.Post.Author).Scan(&postDetail.Author.About,
				&postDetail.Author.Email, &postDetail.Author.FullName, &postDetail.Author.NickName)
		case "thread":
			postDetail.Thread = &Thread{}
			tx.QueryRow(helpers.SelectThreadById, &postDetail.Post.Thread).Scan(
				&postDetail.Thread.ID, &postDetail.Thread.Author, &postDetail.Thread.Created,
				&postDetail.Thread.ForumId, &postDetail.Thread.Message, &postDetail.Thread.Slug,
				&postDetail.Thread.Title, &postDetail.Thread.Votes)
		case "forum":
			postDetail.Forum = &Forum{}

			//_, _ = tx.Exec(helpers.UpdateForumPostsCnt, &postDetail.Post.ForumID)

			tx.QueryRow(helpers.SelectForumDetail, &postDetail.Post.ForumID).Scan(
				&postDetail.Forum.Posts, &postDetail.Forum.Slug, &postDetail.Forum.Threads,
				&postDetail.Forum.Title, &postDetail.Forum.User)
		}
	}
	return &postDetail, nil
}

func (post *Post) PostUpdate(update *PostUpdate, id string) error {
	tx := config.StartTransaction()
	defer tx.Rollback()

	postId, err := strconv.Atoi(id)
	if err != nil {
		return errors.PostNotFound
	}

	err = tx.QueryRow(helpers.SelectPostMessage, &postId).Scan(&post.Author,
		&post.Created, &post.ForumID, &post.ID, &post.IsEdited, &post.Message,
		&post.Parent, &post.Thread)
	if post.Message == update.Message && update.Message != "" {
		return nil
	}

	if err := tx.QueryRow(helpers.UpdatePost, &update.Message, &postId).Scan(&post.Author,
		&post.Created, &post.ForumID, &post.ID, &post.IsEdited, &post.Message,
		&post.Parent, &post.Thread); err != nil {
		return errors.PostNotFound
	}

	tx.Commit()
	return nil
}
