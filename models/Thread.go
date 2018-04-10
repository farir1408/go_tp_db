package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
	"strconv"
	"time"
	"bytes"
	"github.com/jackc/pgx"
)

//easyjson:json
type Thread struct {
	ID      int        `json:"id"`
	Author  string     `json:"author"`
	Created *time.Time `json:"created"`
	ForumId string     `json:"forum"`
	Message string     `json:"message"`
	Slug    string     `json:"slug, omitempty"`
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
	defer tx.Rollback()

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is slug
		if err = tx.QueryRow(helpers.SelectThreadBySlug, slug).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes); err != nil {
			tx.Rollback()
			return errors.ThreadNotFound
		}
	} else {
		if err = tx.QueryRow(helpers.SelectThreadById, id).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes); err != nil {
			tx.Rollback()
			return errors.ThreadNotFound
		}
	}
	return nil
}

func (threadUpdate *ThreadUpdate) ThreadUpdate(slug string) (*Thread, error) {
	tx := config.StartTransaction()

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is string
		row, _ := tx.Exec(helpers.UpdateThreadBySlug, &threadUpdate.Message, &threadUpdate.Title, slug)
		if row.RowsAffected() == 0 {
			tx.Rollback()
			return nil, errors.ThreadNotFound
		}
		thread := Thread{}
		_ = tx.QueryRow(helpers.SelectThreadBySlug, slug).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes)

		tx.Commit()
		return &thread, nil
	} else {
		//slug is id (int)
		row, _ := tx.Exec(helpers.UpdateThreadById, &threadUpdate.Message, &threadUpdate.Title, id)
		if row.RowsAffected() == 0 {
			tx.Rollback()
			return nil, errors.ThreadNotFound
		}
		thread := Thread{}
		_ = tx.QueryRow(helpers.SelectThreadById, slug).Scan(&thread.ID, &thread.Author,
			&thread.Created, &thread.ForumId, &thread.Message, &thread.Slug,
			&thread.Title, &thread.Votes)

		tx.Commit()
		return &thread, nil
	}
}

func GetThreads(slug string, limit []byte,
								since []byte, desc []byte) (Threads, error) {
	tx := config.StartTransaction()
	defer tx.Rollback()
	var results *pgx.Rows
	var err error

	if since != nil {
		if bytes.Equal([]byte("true"), desc) {
			results, err = tx.Query(helpers.SelectThreadsSinceDesc, slug,since, limit)
		} else {
			results, err = tx.Query(helpers.SelectThreadsSince, slug, since, limit)
		}
	} else {
		if bytes.Equal([]byte("true"), desc) {
			results, err = tx.Query(helpers.SelectThreadsDesc, slug, limit)
		} else {
			results, err = tx.Query(helpers.SelectThreads, slug, limit)
		}
	}
	defer results.Close()

	if err != nil {
		log.Fatalln(err)
	}

	threads := Threads{}
	for results.Next() {
		existThread := Thread{}

		if err = results.Scan(&existThread.Author, &existThread.Created, &existThread.ForumId,
			&existThread.ID, &existThread.Message, &existThread.Slug, &existThread.Title);
			err != nil {
				log.Fatalln(err)
			}
		threads = append(threads, &existThread)
	}
	//log.Println(len(threads))
	if len(threads) == 0 {
		var cnt int
		if err = tx.QueryRow("SELECT 1 FROM forum WHERE slug = $1", slug).Scan(&cnt); err != nil {
			return nil, errors.ForumNotFound
		}
	}

	return threads, nil
}
