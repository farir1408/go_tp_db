package models

import (
	"go_tp_db/config"
	"go_tp_db/helpers"
	"log"
)

//easyjson:json
type Status struct {
	ForumSum  int `json:"forum"`
	PostsSum  int `json:"post"`
	ThreadSum int `json:"thread"`
	UserSum   int `json:"user"`
}

func (status *Status) StatusDataBase() {
	tx := config.StartTransaction()
	defer tx.Commit()

	tx.QueryRow(helpers.GetStatus).Scan(&status.ForumSum, &status.PostsSum,
		&status.ThreadSum, &status.UserSum)
}

func ClearDataBase() {
	tx := config.StartTransaction()
	defer tx.Commit()

	_, err := tx.Exec(helpers.ClearDB)
	if err != nil {
		log.Fatalln(err)
	}
}
