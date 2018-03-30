package models

import (
	"time"
	"go_tp_db/config"
	"strconv"
	"go_tp_db/helpers"
	"log"
	"go_tp_db/errors"
)

//easyjson:json
type Thread struct {
	ID      int        `json:"id"`
	Author  string     `json:"author"`
	Created *time.Time `json:"created"`
	ForumId string     `json:"forum"`
	Message string     `json:"message"`
	Slug    string     `json:"slug"`
	Title   string     `json:"title"`
	Votes   int        `json:"votes"`
}

//easyjson:json
type ThreadUpdate struct {
	Message string `json:"message"`
	Title   string `json:"title"`
}

//easyjson:json
type Threads []*Thread

func (thread *Thread) ThreadDetails(slug string) error {
	tx := config.StartTransaction()

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is slug
		if err = tx.QueryRow(helpers.SelectThreadBySlug, slug).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes); err != nil {
				log.Println(err)
				tx.Rollback()
				return errors.ThreadNotFound
		}
	} else {
		if err = tx.QueryRow(helpers.SelectThreadById, id).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes); err != nil {
			log.Println(err)
			tx.Rollback()
			return errors.ThreadNotFound
		}
	}
	return nil
}