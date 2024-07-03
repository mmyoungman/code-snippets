package auth

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"

	"github.com/mmyoungman/sqlitestore"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/openidConnect"
)

func Setup(db *sql.DB) (session *sqlitestore.SqliteStore) {
	sqliteStore, err := sqlitestore.NewSqliteStoreFromConnection(
		db,
		"goth_sessions",
		"/",
		3600,
		[]byte(utils.Getenv("SESSION_SECRET")))
	if err != nil {
		log.Fatal("Failed to create sqlite store", err)
	}
	sqliteStore.MaxLength(8192)
	gothic.Store = sqliteStore

	openidConnect, err := openidConnect.New(
		utils.Getenv("KEYCLOAK_CLIENT_ID"),
		utils.Getenv("KEYCLOAK_CLIENT_SECRET"),
		utils.Getenv("PUBLIC_URL") + "/auth/callback?provider=openid-connect",
		utils.Getenv("KEYCLOAK_DISCOVERY_URL"))
	if err != nil {
		log.Fatal("Error creating openidConnect provider. Error:\n", err)
	}
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	return sqliteStore
}