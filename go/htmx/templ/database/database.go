package database

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	dbFilePath := utils.Getenv("SQLITE3_PATH")

	db, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatal("Failed to open db", err)
	}

	return db
}
