package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// @MarkFix Add Goose for migrations

func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal("Failed to open db", err)
	}

	return db
}