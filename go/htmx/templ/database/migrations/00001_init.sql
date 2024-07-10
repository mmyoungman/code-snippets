-- +goose Up
-- +goose StatementBegin

CREATE TABLE "Sessions" (
	"ID"	TEXT NOT NULL UNIQUE,
	"AccessToken"	TEXT NOT NULL,
	"RefreshToken"	TEXT NOT NULL,
	"TokenType"	TEXT NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE "Sessions";

-- +goose StatementEnd
