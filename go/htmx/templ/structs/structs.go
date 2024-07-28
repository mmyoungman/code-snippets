package structs

import (
	"database/sql"
	"mmyoungman/templ/auth"
)

type ServiceCtx struct { // @MarkFix just make this global?
	Db   *sql.DB
	Auth *auth.Authenticator
}
