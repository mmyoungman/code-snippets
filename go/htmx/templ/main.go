package main

import (
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/handlers"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// include file and line in log messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Didn't load env file", err)
	}

	db := database.DBConnect()
	defer db.Close()

	// @MarkFix the site is currently vulnerable to CSRF attacks?
	authObj, err := auth.Setup()
	if err != nil {
		log.Fatal("Auth setup failed: ", err)
	}

	store.Setup()

	//sessions.NewCookieStore()
	//
	//fileStore := sessions.NewFilesystemStore(
	//	"./tmp",
	//	[]byte(utils.Getenv("SESSION_SECRET")))
	//fileStore.MaxLength(8192)
	//sqliteStore, err := sqlitestore.NewSqliteStoreFromConnection(
	//	db,
	//	"goth_sessions", // @MarkFix Do old sessions ever get removed from this table?
	//	"/",
	//	3600,
	//	[]byte(utils.Getenv("SESSION_SECRET")))
	//if err != nil {
	//	log.Fatal("Failed to create sqlite store", err)
	//}
	//sqliteStore.MaxLength(8192)
	//defer sqliteStore.Close()

	router := chi.NewMux()

	// @MarkFix use other middleware - logger? recoverer?
	// @MarkFix CORS? Use middleware

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
	router.Get("/auth", handlers.Make(handlers.HandleAuthLogin(authObj)))
	router.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback(authObj)))
	router.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout))

	// pages
	router.Get("/", handlers.Make(handlers.HandleHome))

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	port := utils.Getenv("PUBLIC_PORT")
	slog.Info("Starting http server", "URL", utils.Getenv("PUBLIC_HOST")+":"+port)
	if os.Getenv("TEMPL_WATCH_PROXY_URL") != "" {
		slog.Info("Auth configured for watch proxy", "templWatchProxyUrl", utils.Getenv("TEMPL_WATCH_PROXY_URL"))
	}
	// @MarkFix use ListenAndServeTLS
	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
