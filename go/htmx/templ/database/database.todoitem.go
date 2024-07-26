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

func GetToDoItem(db *sql.DB, toDoItemID string) *model.ToDoItem {
	stmt := SELECT(ToDoItem.AllColumns).
		FROM(ToDoItem).
		WHERE(ToDoItem.ID.EQ(String(toDoItemID)))

	var items []model.ToDoItem
	err := stmt.Query(db, &items)
	if err != nil {
		log.Fatal("Failed to execute SQL query", err)
	}

	if len(items) == 0 {
		return nil
	}

	return &items[0]
}

func InsertToDoItem(db *sql.DB, item *model.ToDoItem) *model.ToDoItem {
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

	return item
}

func UpdateToDoItem(db *sql.DB, item *model.ToDoItem) *model.ToDoItem {
	stmt := ToDoItem.
		UPDATE(ToDoItem.Name, ToDoItem.Description).
		SET(item.Name, item.Description).
		WHERE(ToDoItem.ID.EQ(String(item.ID)))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one ToDoItem")
	}

	return item
}

func DeleteToDoItem(db *sql.DB, id string) {
	stmt := ToDoItem.
		DELETE().
		WHERE(ToDoItem.ID.EQ(String(id)))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have removed one ToDoItem")
	}
}
