package database

import (
	"database/sql"
	"log"

	"mmyoungman/templ/database/jet/model"
	. "mmyoungman/templ/database/jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func ListToDoItem(db *sql.DB, toDoID string) []*model.ToDoItem {
	stmt := SELECT(User.AllColumns).
		FROM(ToDoItem).
		WHERE(User.ID.EQ(String(toDoID)))

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
		log.Fatal("Should have added one User")
	}
}

//func InsertUser(db *sql.DB, user *model.User) {
//	stmt := User.
//		INSERT(User.ID, User.Username, User.Email, User.FirstName, User.LastName, User.RawIDToken).
//		VALUES(user.ID, user.Username, user.Email, user.FirstName, user.LastName, user.RawIDToken)
//
//	result, err := stmt.Exec(db)
//	if err != nil {
//		log.Fatal("Failed to execute query ", err)
//	}
//
//	n, _ := result.RowsAffected()
//	if n != 1 {
//		log.Fatal("Should have added one User")
//	}
//}
//
//func UpdateUser(db *sql.DB, user *model.User) {
//	stmt := User.
//		UPDATE(User.Username, User.Email, User.FirstName, User.LastName, User.RawIDToken).
//		SET(user.Username, user.Email, user.FirstName, user.LastName, user.RawIDToken).
//		WHERE(User.ID.EQ(String(user.ID)))
//
//	result, err := stmt.Exec(db)
//	if err != nil {
//		log.Fatal("Failed to execute query ", err)
//	}
//
//	n, _ := result.RowsAffected()
//	if n != 1 {
//		log.Fatal("Should have updated one Session")
//	}
//}

//func DeleteUser(db *sql.DB, userID string) {
//	stmt := User.DELETE().
//		WHERE(User.ID.EQ(String(userID)))
//
//	result, err := stmt.Exec(db)
//	if err != nil {
//		log.Fatal("Failed to execute query ", err)
//	}
//
//	n, _ := result.RowsAffected()
//	if n != 1 {
//		log.Fatal("Should have updated one Session")
//	}
//}
