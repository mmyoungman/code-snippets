package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func CreateDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		return nil, err
	}
	// NOTE: function caller handles Close

	{
		var version string
		err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
		if err != nil {
			return nil, err
		}
	}

	query := `
  CREATE TABLE IF NOT EXISTS Users(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL
  );
  `
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	query = `
  INSERT OR IGNORE INTO Users (email, password)
  VALUES ('user@test.com', 'password');
  `
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}
