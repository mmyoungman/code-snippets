-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS "Session" (
	"ID"	TEXT NOT NULL UNIQUE,
	"UserID"	TEXT NOT NULL,
	"AccessToken"	TEXT NOT NULL,
	"RefreshToken"	TEXT NOT NULL,
    "Expiry" INT NOT NULL,
	"TokenType"	TEXT NOT NULL,
	PRIMARY KEY("ID")
);

CREATE TABLE IF NOT EXISTS "User" (
	"ID" TEXT NOT NULL UNIQUE,
	"Username" TEXT NOT NULL,
	"Email" TEXT NOT NULL,
	"FirstName" TEXT NOT NULL,
	"LastName" TEXT NOT NULL,
	"RawIDToken" TEXT NOT NULL,
	PRIMARY KEY("ID")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS "Session";
DROP TABLE IF EXISTS "User";

-- +goose StatementEnd
