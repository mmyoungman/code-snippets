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

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "Session";

-- +goose StatementEnd
