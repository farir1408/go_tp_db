package models

import (
	"bytes"
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
	defer tx.Rollback()

	rows, err := tx.Exec(helpers.CreateUser, &user.About, &user.Email, &user.FullName, &user.NickName)
	if err != nil {
		log.Panic(err)
	}

	//rows != 0 if user was created in Exec command, if user was created earlier rows = 0
	if rows.RowsAffected() == 0 {
		userArr := Users{}
		queryRows, err := tx.Query(helpers.SelectUser, &user.NickName, &user.Email)
		if err != nil {
			log.Panic(err)
		}

		defer queryRows.Close()

		for queryRows.Next() {
			isUserExist := User{}
			queryRows.Scan(&isUserExist.About, &isUserExist.Email,
				&isUserExist.FullName, &isUserExist.NickName)
			userArr = append(userArr, &isUserExist)
		}

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
		return errors.UserNotFound
	}

	return nil
}

func (newUser *User) UpdateUserProfile() error {
	tx := config.StartTransaction()
	defer tx.Rollback()

	if err := tx.QueryRow(helpers.UpdateUser, &newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName).Scan(&newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName); err != nil {
		if _, ok := err.(pgx.PgError); ok {
			return errors.UserUpdateConflict
		}

		return errors.UserNotFound
	}
	if _, err := tx.Exec(helpers.UpdateForumUsers, &newUser.About, &newUser.Email,
		&newUser.FullName, &newUser.NickName); err != nil {
		log.Panic(err)
	}

	tx.Commit()
	return nil
}

func GetUsers(slug string, limit []byte, since []byte,
	desc []byte) (Users, error) {

	tx := config.StartTransaction()
	defer tx.Rollback()
	var results *pgx.Rows
	var err error

	if since != nil {
		if bytes.Equal([]byte("true"), desc) {
			results, err = tx.Query(helpers.SelectUsersSinceDesc, slug, string(since), limit)
		} else {
			results, err = tx.Query(helpers.SelectUsersSince, slug, string(since), limit)
		}
	} else {
		if bytes.Equal([]byte("true"), desc) {
			results, err = tx.Query(helpers.SelectUsersDesc, slug, limit)
		} else {
			results, err = tx.Query(helpers.SelectUsers, slug, limit)
		}
	}

	defer results.Close()

	users := Users{}
	for results.Next() {
		user := User{}

		if err = results.Scan(&user.About, &user.Email, &user.FullName, &user.NickName); err != nil {
			log.Fatalln(err)
		}
		users = append(users, &user)
	}

	if len(users) == 0 {
		var cnt int
		if err = tx.QueryRow("SELECT 1 FROM forum WHERE slug = $1", slug).Scan(&cnt); err != nil {
			return nil, errors.ForumNotFound
		}
	}

	return users, nil
}
