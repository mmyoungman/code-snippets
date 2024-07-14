package database

import (
	"database/sql"
	"log"

	"mmyoungman/templ/database/jet/model"
	. "mmyoungman/templ/database/jet/table"

	. "github.com/go-jet/jet/v2/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func GetSession(db *sql.DB, sessionID string) *model.Session {
	stmt := SELECT(Session.AllColumns).
		FROM(Session).
		WHERE(Session.ID.EQ(String(sessionID)))

	var sessions []model.Session
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
	stmt := Session.
		INSERT(Session.ID, Session.UserID, Session.AccessToken, Session.RefreshToken, Session.Expiry, Session.TokenType).
		VALUES(sessionID, userID, accessToken, refreshToken, expiry, tokenType)

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
	stmt := Session.
		UPDATE(Session.AccessToken, Session.RefreshToken, Session.Expiry, Session.TokenType).
		SET(accessToken, refreshToken, expiry, tokenType).
		WHERE(Session.ID.EQ(String(sessionID)).
			AND(Session.UserID.EQ(String(userID))))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one Session")
	}
}

func DeleteSession(db *sql.DB, sessionID string) {
	stmt := Session.DELETE().
		WHERE(Session.ID.EQ(String(sessionID)))

	result, err := stmt.Exec(db)
	if err != nil {
		log.Fatal("Failed to execute query ", err)
	}

	n, _ := result.RowsAffected()
	if n != 1 {
		log.Fatal("Should have updated one Session")
	}
}
