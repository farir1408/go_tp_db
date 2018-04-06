package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"strconv"
	"time"
)

//easyjson:json
type Post struct {
	ID       int        `json:"id"`
	Author   string     `json:"author"`
	Created  *time.Time `json:"created"`
	ForumID  string     `json:"forum"`
	IsEdited bool       `json:"isEdited"`
	Message  string     `json:"message"`
	Parent   int        `json:"parent, omitempty"`
	Thread   int        `json:"thread"`
	Slug	string		`json:"slug, omitempty"`
}

//easyjson:json
type PostDetail struct {
	Author *User   `json:"author"`
	Forum  *Forum  `json:"forum"`
	Post   *Post   `json:"post"`
	Thread *Thread `json:"thread"`
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

	if len(*posts) == 0 {
		return errors.NoPostsForCreate
	}

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
	//check
	currentTime := time.Now().Format("2000-01-01T00:00:00.000Z")
	curTime := time.Now()

	for _, post := range *posts {
		var parentId int
		if post.Parent != 0 {
			err = tx.QueryRow(helpers.SelectThreadID, &post.Parent).Scan(&parentId)
			if err != nil {
				return errors.NoThreadParent
			}
		}

		if err = tx.QueryRow(helpers.CreatePost, post.Author, currentTime,
			forumSlug, post.Message, parentId, id).Scan(&post.ID); err != nil {
				return errors.NoThreadParent
		}
		post.Created = &curTime
		post.IsEdited = false
		post.ForumID = forumSlug
		post.Thread = id

	}
	tx.Commit()
	return nil

	//TODO: complete the request processing
}
