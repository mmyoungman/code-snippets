package main

import (
	"database/sql"
	"fmt"
	"log"
	"mmyoungman/nostr_backup/pkg/json"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./nostr_backup.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS Events(
		id TEXT UNIQUE,
		pubkey TEXT,
		created_at UNSIGNED INT(2),
		kind int,
		tags TEXT,
		content TEXT,
		sig TEXT
	);`

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DBInsertEvent(db *sql.DB, event Event) {
	query := `
	SELECT count(1) FROM Events
	WHERE id = ?;`

	result, err := db.Query(query, event.Id)
	var exists bool
	result.Next()
	result.Scan(&exists)
	result.Close()
	if exists {
		fmt.Println(event.Id, "already in DB")
		return
	}

	query = `
	INSERT INTO Events (id, pubkey, created_at, kind, tags, content, sig)
	VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err = db.Exec(query, event.Id,
		event.PubKey, event.CreatedAt, event.Kind,
		event.Tags.ToJson(), DecorateJsonStr(event.Content), event.Sig)
	if err != nil {
		log.Fatal(err)
	}

}

func DBGetEvents(db *sql.DB) []Event {
	query := `
	SELECT id, pubkey, created_at, kind, tags, content, sig
	FROM Events;`

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var events []Event = make([]Event, 0)
	for rows.Next() {
		var event Event
		var tags string

		err = rows.Scan(&event.Id, &event.PubKey, &event.CreatedAt,
		&event.Kind, &tags, &event.Content, &event.Sig)
		if err != nil {
			log.Fatal(err)
		}
		json.UnmarshalJSON([]byte(tags), event.Tags)

		events = append(events, event)
	}

	return events
}
