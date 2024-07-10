-- +goose Up
-- +goose StatementBegin

CREATE TABLE "Sessions" (
	"ID"	TEXT NOT NULL UNIQUE,
	"UserID"	TEXT NOT NULL,
	"AccessToken"	TEXT NOT NULL,
	"RefreshToken"	TEXT NOT NULL,
	"TokenType"	TEXT NOT NULL,
	PRIMARY KEY("ID")
)

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "Sessions";

-- +goose StatementEnd
