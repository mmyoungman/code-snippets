package main

import (
	"database/sql"
	"fmt"
	"log"
	"mmyoungman/nostr_backup/internal/json"

	_ "github.com/mattn/go-sqlite3"
)

func DBConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "./nostr_backup.db")
	if err != nil {
		log.Fatal(err)
	}

	stm, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS Events(
		id TEXT UNIQUE,
		pubkey TEXT,
		created_at UNSIGNED INT(2),
		kind INT,
		tags TEXT,
		content TEXT,
		sig TEXT
	);`)
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()

	_, err = stm.Exec()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DBInsertEvent(db *sql.DB, event Event) {
	stm, err := db.Prepare(`
	SELECT count(1) FROM Events
	WHERE id = ?;`)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stm.Query(event.Id)
	var exists bool
	result.Next()
	result.Scan(&exists)
	result.Close()
	if exists {
		fmt.Println(event.Id, "already in DB")
		return
	}
	stm.Close()

	stm, err = db.Prepare(`
	INSERT INTO Events (id, pubkey, created_at, kind, tags, content, sig)
	VALUES (?, ?, ?, ?, ?, ?, ?)`)

	_, err = stm.Exec(event.Id,
		event.PubKey, event.CreatedAt, event.Kind,
		event.Tags.ToJson(), DecorateJsonStr(event.Content), event.Sig)
	if err != nil {
		log.Fatal(err)
	}

}

func DBGetEvents(db *sql.DB) []Event {
	stm, err := db.Prepare(`
	SELECT id, pubkey, created_at, kind, tags, content, sig
	FROM Events;`)
	if err != nil {
		log.Fatal(err)
	}
	defer stm.Close()

	rows, err := stm.Query()
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
