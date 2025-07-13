-- +goose Up
-- +goose StatementBegin

CREATE TABLE Sessions (
	ID TEXT NOT NULL UNIQUE,
	Userid TEXT NOT NULL,
	Accesstoken TEXT NOT NULL,
	Refreshtoken TEXT NOT NULL,
    Expiry INT NOT NULL,
	Tokentype TEXT NOT NULL,
	PRIMARY KEY(ID)
);

CREATE TABLE Users (
	ID TEXT NOT NULL UNIQUE,
	Username TEXT NOT NULL,
	Email TEXT NOT NULL,
	Firstname TEXT NOT NULL,
	Lastname TEXT NOT NULL,
	RawIdtoken TEXT NOT NULL,
	PRIMARY KEY(ID)
);

CREATE TABLE ToDoItems (
	ID TEXT NOT NULL UNIQUE,
	Name TEXT NOT NULL,
	Description TEXT NOT NULL,
	PRIMARY KEY(ID)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS Sessions;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS ToDoItems;

-- +goose StatementEnd
