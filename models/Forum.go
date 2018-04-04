package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
	"strings"
)

//easyjson:json
type Forum struct {
	Posts   int64  `json:"posts"`
	Slug    string `json:"slug"`
	Threads int32  `json:"threads"`
	Title   string `json:"title"`
	User    string `json:"user"`
}

//easyjson:json
type ForumDetail struct {
	Slug string `json:"slug"`
}

func (forum *Forum) ForumCreate() (*Forum, error) {
	//start database transaction
	tx := config.StartTransaction()
	forum.User = strings.ToLower(forum.User)
	log.Println(forum.User)
	isForumExist := Forum{}

	//checking the forum for existence
	if err := tx.QueryRow(helpers.SelectForumCreate, &forum.Slug,
		&forum.User).Scan(&isForumExist.Posts,
		&isForumExist.Slug, &isForumExist.Threads,
		&isForumExist.Title, &isForumExist.User); err == nil {
		tx.Rollback()
		return &isForumExist, errors.ForumIsExist
	}

	//checking the author(user) for existence
	if err := tx.QueryRow(helpers.SelectForumByUser, &forum.Slug, &forum.Title, &forum.User).Scan(&forum.User); err != nil {
		//log.Println(err)
		tx.Rollback()
		return nil, errors.UserNotFound
	}
	log.Println("WARNING")
	log.Println(string(forum.User))

	tx.Commit()
	return nil, nil
}

func (forum *Forum) ForumDetails(slug string) error {
	//start database transaction
	tx := config.StartTransaction()

	if err := tx.QueryRow(helpers.SelectForumDetail, slug).Scan(&forum.Posts, &forum.Slug,
		&forum.Threads, &forum.Title, &forum.User); err != nil {
		tx.Rollback()
		return errors.UserNotFound
	}

	return nil
}

func (thread *Thread) ForumThreadCreate() (*Thread, error) {
	//start database transaction
	tx := config.StartTransaction()
	log.Println("SLUG is - ", thread.Slug)

	isThreadExist := Thread{}

	if err := tx.QueryRow(helpers.SelectThreadCreate, &thread.Message, &thread.Title,
		&thread.Slug).Scan(
		&isThreadExist.Author, &isThreadExist.Created, &isThreadExist.ForumId, &isThreadExist.ID,
		&isThreadExist.Message, &isThreadExist.Slug, &isThreadExist.Title, &isThreadExist.Votes); err == nil {
		tx.Rollback()
		return &isThreadExist, errors.ThreadIsExist
	}

	if err := tx.QueryRow(helpers.SelectThreadByUser, &thread.Author, &thread.Message,
		&thread.Title, &thread.ForumId, &thread.Slug, &thread.Created).Scan(&thread.ID); err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, errors.ForumNotFound
	}

	tx.Commit()
	return nil, nil
}
