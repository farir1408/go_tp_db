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

	id, err := strconv.Atoi(slug)

	if err != nil {
		//slug is slug (string)\
		if _ = tx.QueryRow(helpers.CreateVoteIdBySlug, slug).Scan(&id); id == 0 {
			return errors.ThreadNotFound
		}
	}
	row, err := tx.Exec(helpers.UpdateVoteById, &vote.Voice, id, &vote.NickName)
	if err != nil {
		return errors.ThreadNotFound
	}

	if row.RowsAffected() == 0 {
		if err = tx.QueryRow(helpers.SelectThreadId, id).Scan(&id); err != nil {
			return errors.ThreadNotFound
		}

		row, err = tx.Exec(helpers.CreateVoteById, &vote.Voice, &vote.NickName, id)
		if err != nil {
			return errors.ThreadNotFound
		}
	}
	_, _ = tx.Exec(helpers.UpdateThreadVotes, id)
	if row.RowsAffected() != 0 {
		tx.Commit()
		return nil
	}
	return nil
}