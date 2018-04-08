package config

import (
	"github.com/jackc/pgx"
	"io/ioutil"
	"log"
)

var db *pgx.ConnPool

var pgxConfig = pgx.ConnConfig{
	Host:     "localhost",
	Port:     5432,
	Database: "tpdb",
	User:     "postgres",
	Password: "postgres",
}

const dataBaseSchema = "./config/database_schema.sql"

func loadSchemaSQL() error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(err)
		return err
	}

	content, err := ioutil.ReadFile(dataBaseSchema)
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err = tx.Exec(string(content)); err != nil {
		log.Println(err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func StartTransaction() *pgx.Tx {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	return tx
}

func InitDB() {
	log.Println("start initDB")
	var err error

	db, err = pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: 100,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err = loadSchemaSQL(); err != nil {
		log.Println("Error is exist", err)
	}
	log.Println("schema initialized successfull")
}
