package models

import (
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"strconv"
)

//easyjson:json
type Vote struct {
	NickName string `json:"nickname"`
	Voice    int    `json:"voice"`
}

func (vote *Vote) Vote(slug string) error {
	tx := config.StartTransaction()
	defer tx.Rollback()

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is slug (string)
		if _ = tx.QueryRow(helpers.CreateVoteIdBySlug, slug).Scan(&id); id == 0 {
			return errors.ThreadNotFound
		}
	}

	var diff int
	tx.QueryRow(helpers.UpdateVoteById, &vote.Voice, id, &vote.NickName).Scan(&diff)

	if diff == 0 {

		row, err := tx.Exec(helpers.CreateVoteById, &vote.Voice, &vote.NickName, id)
		if err != nil {
			return errors.ThreadNotFound
		}

		if row.RowsAffected() != 0 {
			diff = vote.Voice
		}
	}
	_, _ = tx.Exec(helpers.UpdateThreadVotes, &diff, &id)


	tx.Commit()
	return nil
}
