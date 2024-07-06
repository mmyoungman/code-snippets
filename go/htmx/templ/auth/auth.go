package auth

import (
	"database/sql"
	"log"
	"mmyoungman/templ/utils"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/openidConnect"
	"github.com/mmyoungman/sqlitestore"
)

func Setup(db *sql.DB) (session *sessions.Store) {
	var store sessions.Store

	if (utils.IsDev) { // @MarkFix Do I need to bother with this?
		fileStore := sessions.NewFilesystemStore(
			"./tmp",
			[]byte(utils.Getenv("SESSION_SECRET")))
		fileStore.MaxLength(8192)

		gothic.Store = fileStore
		store = sessions.Store(fileStore)
	} else {
		sqliteStore, err := sqlitestore.NewSqliteStoreFromConnection(
			db,
			"goth_sessions", // @MarkFix Do old sessions ever get removed from this table?
			"/",
			3600,
			[]byte(utils.Getenv("SESSION_SECRET")))
		if err != nil {
			log.Fatal("Failed to create sqlite store", err)
		}
		sqliteStore.MaxLength(8192)
		gothic.Store = sqliteStore
		store = sessions.Store(sqliteStore)
	}

	var callbackHost string
	// Configure auth to work with templ watch proxy if we're using it @hotreload
	templProxyURL := os.Getenv("TEMPL_WATCH_PROXY_URL") // use os.Getenv here becuase we don't want to trigger a log.Fatal
	if templProxyURL == "" {
		callbackHost = utils.Getenv("PUBLIC_HOST") + ":" + utils.Getenv("PUBLIC_PORT")
	} else {
		callbackHost = templProxyURL
	}

	keycloakOIDC, err := openidConnect.New(
		utils.Getenv("KEYCLOAK_CLIENT_ID"),
		utils.Getenv("KEYCLOAK_CLIENT_SECRET"),
		callbackHost + "/auth/callback?provider=openid-connect",
		utils.Getenv("KEYCLOAK_CALLBACK_URL"))
	if err != nil {
		log.Fatal("Error creating openidConnect provider. Error:\n", err)
	}
	if keycloakOIDC != nil {
		goth.UseProviders(keycloakOIDC)
	}

	return &store
}