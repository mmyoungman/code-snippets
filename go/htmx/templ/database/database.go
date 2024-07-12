package database

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"

	"mmyoungman/templ/database/jet/model"
	. "mmyoungman/templ/database/jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
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

func GetSession(db *sql.DB, sessionID string, userID string) *model.Sessions {
	stmt := SELECT(Sessions.AllColumns).FROM(
			Sessions).WHERE(
				Sessions.ID.EQ(String(sessionID),
				).AND(
					Sessions.UserID.EQ(String(userID)),
				),
			)
	
	var sessions []model.Sessions
	err := stmt.Query(db, &sessions)
	if err != nil {
		log.Fatal("Failed to execute SQL query", err)
	}

	if len(sessions) == 0 {
		// @MarkFix shouldn't have > 1 ever - should error on that
		return nil 
	}

	return &sessions[0]
}

func InsertSession(db *sql.DB, sessionID string, userID string, accessToken string, refreshToken string, expiry int64, tokenType string) {
	stmt := Sessions.INSERT(
		Sessions.ID, Sessions.UserID, Sessions.AccessToken, Sessions.RefreshToken, Sessions.Expiry, Sessions.TokenType).VALUES(
			sessionID, userID, accessToken, refreshToken, expiry, tokenType)
	
	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have added one Session")
	}
}

func UpdateSession(db *sql.DB, sessionID string, userID string, accessToken string, refreshToken string, expiry int64, tokenType string) {
	stmt := Sessions.UPDATE(
			Sessions.AccessToken, Sessions.RefreshToken, Sessions.Expiry, Sessions.TokenType,
		).SET(
			accessToken, refreshToken, expiry, tokenType,	
		).WHERE(
			Sessions.ID.EQ(String(sessionID),
			).AND(Sessions.UserID.EQ(String(userID))),
		)
	
	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	//queryStr, args := stmt.Sql()

	//queryStmt, err := db.Prepare(queryStr)
	//if err != nil {
	//	log.Fatal("Invalid SQL query", err)
	//}
	//defer queryStmt.Close()

	//result, err := queryStmt.Exec(args)
	//if err != nil {
	//	log.Fatal("Failed to execute SQL query", err)
	//}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one Session")
	}
}

func DeleteSession(db *sql.DB, sessionID string) {
	stmt := Sessions.DELETE().WHERE(Sessions.ID.EQ(String(sessionID)))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one Session")
	}
}