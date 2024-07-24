package database

import (
	"database/sql"
	"log"

	"mmyoungman/templ/database/jet/model"
	. "mmyoungman/templ/database/jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func ListToDoItems(db *sql.DB) []*model.ToDoItem {
	stmt := SELECT(ToDoItem.AllColumns).
		FROM(ToDoItem)

	var items []*model.ToDoItem
	err := stmt.Query(db, &items)
	if err != nil {
		log.Fatal("Failed to execute SQL query", err)
	}

	return items
}

func InsertToDoItem(db *sql.DB, item *model.ToDoItem) {
	stmt := ToDoItem.
		INSERT(ToDoItem.ID, ToDoItem.Name, ToDoItem.Description).
		VALUES(item.ID, item.Name, item.Description)
	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have added one ToDoItem")
	}
}
