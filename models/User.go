package models

import (
	"log"
	"go_tp_db/config"
	"go_tp_db/helpers"
	"go_tp_db/errors"
)

//easyjson:json
type User struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	NickName string `json:"nickname, omitempty"`
}

//easyjson:json
type UserUpdate struct {
	About    string `json:"about"`
	Email    string `json:"email"`
	FullName string `json:"fullname"`
}

//easyjson:json
type Users []*User

func (user *User) UserCreate() (Users, error) {
	tx := config.StartTransaction()

	rows, err := tx.Exec(helpers.CreateUser, &user.About, &user.Email, &user.FullName, &user.NickName)
	if err != nil {
		log.Println("WARNING", err)
	}

	//rows != 0 if user was created in Exec command, if user was created earlier rows = 0
	if rows.RowsAffected() == 0 {
		userArr := Users{}
		queryRows, err := tx.Query(helpers.SelectUser, &user.NickName, &user.Email)
		if err != nil {
			log.Println(err)
		}

		defer queryRows.Close()

		for queryRows.Next() {
			isUserExist := User{}
			log.Println("We are hear!!!")
			queryRows.Scan(&isUserExist.About, &isUserExist.Email,
				&isUserExist.FullName, &isUserExist.NickName)

			log.Println(&isUserExist.NickName, &isUserExist.FullName)
			userArr = append(userArr, &isUserExist)
		}

		tx.Rollback()
		return userArr, errors.UserIsExist
	}

	tx.Commit()
	return nil, nil
}

func (user *User) UserProfile(nickname string) error {
	tx := config.StartTransaction()

	if err := tx.QueryRow(helpers.SelectUserProfile, nickname).Scan(&user.About,
		&user.Email, &user.FullName, &user.NickName); err != nil {
		tx.Rollback()
		return errors.UserNotFound
	}

	return nil
}

func (newUser *User) UpdateUserProfile() error {
	log.Println("somthing")
	tx := config.StartTransaction()

	rows, err := tx.Exec(helpers.UpdateUser, &newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return errors.UserUpdateConflict
	}

	if rows.RowsAffected() == 0 {
		log.Println(err)
		tx.Rollback()
		return errors.UserNotFound
	}
	tx.Commit()
	log.Println(err)
	return nil
}