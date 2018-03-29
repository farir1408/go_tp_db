package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
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

	isForumExist := Forum{}

	//checking the forum for existence
	if err := tx.QueryRow(helpers.SelectForumCreate, &forum.Slug,
		&forum.Title, &forum.User).Scan(&isForumExist.Posts,
		&isForumExist.Slug, &isForumExist.Threads,
		&isForumExist.Title, &isForumExist.User); err == nil {
		tx.Rollback()
		return &isForumExist, errors.ForumIsExist
	}

	//checking the author(user) for existence
	if _, err := tx.Exec(helpers.SelectForumByUser, &forum.Slug, &forum.Title, &forum.User); err != nil {
		//log.Println(err)
		tx.Rollback()
		return nil, errors.UserNotFound
	}

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

	isThreadExist := Thread{}

	if err := tx.QueryRow(helpers.SelectThreadCreate, &thread.Author, &thread.Message, &thread.Title,
		&thread.Created).Scan(
		&isThreadExist.Author, &isThreadExist.Created, &isThreadExist.ForumId, &isThreadExist.ID,
		&isThreadExist.Message, &isThreadExist.Slug, &isThreadExist.Title, &isThreadExist.Votes); err == nil {
		tx.Rollback()
		return &isThreadExist, errors.ThreadIsExist
	}

	if _, err := tx.Exec(helpers.SelectThreadByUser, &thread.Author, &thread.Message,
		&thread.Title, &thread.ForumId, &thread.Slug, &thread.Created); err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, errors.ForumNotFound
	}

	tx.Commit()
	return nil, nil
}
