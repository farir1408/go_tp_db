package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
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