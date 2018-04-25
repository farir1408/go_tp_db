package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
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
	defer tx.Rollback()
	forum.User = strings.ToLower(forum.User)
	isForumExist := Forum{}

	//checking the forum for existence
	if err := tx.QueryRow(helpers.SelectForumCreate, &forum.Slug,
		&forum.User).Scan(&isForumExist.Posts,
		&isForumExist.Slug, &isForumExist.Threads,
		&isForumExist.Title, &isForumExist.User); err == nil {
		return &isForumExist, errors.ForumIsExist
	}

	//checking the author(user) for existence
	if err := tx.QueryRow(helpers.InsertForumByUser, &forum.Slug, &forum.Title, &forum.User).Scan(&forum.User); err != nil {
		return nil, errors.UserNotFound
	}

	tx.Commit()
	return nil, nil
}

func (forum *Forum) ForumDetails(slug string) error {
	//start database transaction
	tx := config.StartTransaction()
	defer tx.Rollback()

	if err := tx.QueryRow(helpers.SelectForumDetail, slug).Scan(&forum.Posts, &forum.Slug,
		&forum.Threads, &forum.Title, &forum.User); err != nil {
		return errors.UserNotFound
	}

	return nil
}

func (thread *Thread) ForumThreadCreate() (*Thread, error) {
	//start database transaction
	tx := config.StartTransaction()
	defer tx.Rollback()

	isThreadExist := Thread{}
	if thread.Slug == "" {
		if err := tx.QueryRow(helpers.SelectThreadCreate, &thread.Message, &thread.Title).Scan(
			&isThreadExist.Author, &isThreadExist.Created, &isThreadExist.ForumId, &isThreadExist.ID,
			&isThreadExist.Message, &isThreadExist.Slug, &isThreadExist.Title, &isThreadExist.Votes); err == nil {
			return &isThreadExist, errors.ThreadIsExist
		}
	} else {
		if err := tx.QueryRow(helpers.SelectThreadCreateSlug, &thread.Slug).Scan(
			&isThreadExist.Author, &isThreadExist.Created, &isThreadExist.ForumId, &isThreadExist.ID,
			&isThreadExist.Message, &isThreadExist.Slug, &isThreadExist.Title, &isThreadExist.Votes); err == nil {
			return &isThreadExist, errors.ThreadIsExist
		}
	}

	if err := tx.QueryRow(helpers.SelectThreadByUser, &thread.Author, &thread.Message,
		&thread.Title, &thread.ForumId, &thread.Slug, &thread.Created).Scan(&thread.ID,
		&thread.ForumId); err != nil {
		return nil, errors.ForumNotFound
	}

	tx.Exec(helpers.InsertForumUsersTmpThread, &thread.Author, &thread.ForumId)
	tx.Exec(helpers.UpdateThreadCntForum, &thread.ForumId)

	tx.Commit()
	return nil, nil
}
