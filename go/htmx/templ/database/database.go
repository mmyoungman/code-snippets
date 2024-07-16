package database

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"

	_ "github.com/mattn/go-sqlite3"
)

// @MarkFix swap go-jet for sqlc
// @MarkFix sqlite tuning: https://kerkour.com/sqlite-for-servers
// https://www.reddit.com/r/golang/comments/1e4m07d/using_mutex_while_writing_to_sqlite_database/

func Connect() *sql.DB {
	dbFilePath := utils.Getenv("SQLITE3_PATH")

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal("Failed to open db", err)
	}

	return db
}
