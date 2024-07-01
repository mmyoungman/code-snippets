package database

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"

	_ "github.com/mattn/go-sqlite3"
)

// @MarkFix Add goose for migrations
// @MarkFix Use go-jet as SQL builder

func DBConnect() *sql.DB {
	dbPath := utils.Getenv("SQLITE3_PATH")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal("Failed to open db", err)
	}

	return db
}