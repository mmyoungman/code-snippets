package database

import (
	"database/sql"
	"log"

	"mmyoungman/templ/database/jet/model"
	. "mmyoungman/templ/database/jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func GetUser(db *sql.DB, userID string) *model.User {
	stmt := SELECT(User.AllColumns).
		FROM(User).
		WHERE(User.ID.EQ(String(userID)))

	var users []model.User
	err := stmt.Query(db, &users)
	if err != nil {
		log.Fatal("Failed to execute SQL query", err)
	}

	if len(users) != 1 {
		return nil
	}

	return &users[0]
}

func InsertUser(db *sql.DB, user *model.User) {
	stmt := User.
		INSERT(User.ID, User.Username, User.Email, User.FirstName, User.LastName, User.RawIDToken).
		VALUES(user.ID, user.Username, user.Email, user.FirstName, user.LastName, user.RawIDToken)

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have added one User")
	}
}

func UpdateUser(db *sql.DB, user *model.User) {
	stmt := User.
		UPDATE(User.Username, User.Email, User.FirstName, User.LastName, User.RawIDToken).
		SET(user.Username, user.Email, user.FirstName, user.LastName, user.RawIDToken).
		WHERE(User.ID.EQ(String(user.ID)))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one Session")
	}
}

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
