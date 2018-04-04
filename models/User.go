package models

import (
	"github.com/jackc/pgx"
	"go_tp_db/config"
	"go_tp_db/errors"
	"go_tp_db/helpers"
	"log"
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
	About    string `json:"about, omitempty"`
	Email    string `json:"email, omitempty"`
	FullName string `json:"fullname, omitempty"`
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
			//log.Println("New user")
			queryRows.Scan(&isUserExist.About, &isUserExist.Email,
				&isUserExist.FullName, &isUserExist.NickName)

			//log.Println(&isUserExist.NickName, &isUserExist.FullName)
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
	defer tx.Rollback()

	if err := tx.QueryRow(helpers.SelectUserProfile, nickname).Scan(&user.About,
		&user.Email, &user.FullName, &user.NickName); err != nil {
		log.Println(err)
		tx.Rollback()
		return errors.UserNotFound
	}

	return nil
}

func (newUser *User) UpdateUserProfile() error {
	tx := config.StartTransaction()

	if err := tx.QueryRow(helpers.UpdateUser, &newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName).Scan(&newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName); err != nil {
		if _, ok := err.(pgx.PgError); ok {
			tx.Rollback()
			return errors.UserUpdateConflict
		}
		tx.Rollback()
		return errors.UserNotFound
	}

	tx.Commit()
	return nil

	//if err != nil {
	//	//log.Println("CONFLICT", err)
	//	tx.Rollback()
	//	return errors.UserUpdateConflict
	//}
	//
	//if rows.RowsAffected() == 0 {
	//	//log.Println("NotUSER", err)
	//	tx.Rollback()
	//	return errors.UserNotFound
	//}

	tx.Commit()
	return nil
}
