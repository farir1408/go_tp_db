package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
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
	Thread   string     `json:"thread"`
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

	//checking thread id or slug
	var forum string

	id, err := strconv.Atoi(slug)
	if err != nil {
		//	id is slug (string)
		if err = tx.QueryRow(helpers.SelectThreadIdForumId, slug).Scan(&id, &forum); err != nil {
			log.Println(err)
			return errors.ThreadNotFound
		}
	} else {
		if err = tx.QueryRow(helpers.SelectThreadIdForumIdByID, id).Scan(&id, &forum); err != nil {
			log.Println(err)
			return errors.ThreadNotFound
		}
	}
	//check
	log.Println(forum)
	log.Println("thread id ", id)
	log.Println("thread slug ", slug)
	return nil

	//TODO: complete the request processing
}
