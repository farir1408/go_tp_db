package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
	"strconv"
)

//easyjson:json
type Vote struct {
	NickName string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func (vote *Vote) Vote(slug string) error {
	tx := config.StartTransaction()

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is slug (string)
		log.Println("PASS")
		return errors.ThreadNotFound
	} else {
		row, err := tx.Exec(helpers.CreateVoteById, &vote.Voice, &vote.NickName, id)
		if err != nil {
			tx.Rollback()
			log.Println(err)
			return errors.ThreadNotFound
		}
		if row.RowsAffected() != 0 {
			tx.Commit()
			return nil
		}
	}
	return nil
}
