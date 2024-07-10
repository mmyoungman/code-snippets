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

	// @MarkFix make the program print all logs etc. to a file

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Didn't load env file", err)
	}

	db := database.DBConnect()
	defer db.Close()

	authObj, err := auth.Setup()
	if err != nil {
		log.Fatal("Auth setup failed: ", err)
	}

	store.Setup()

	// @MarkFix build pipeline?
	// @MarkFix I suppose I could write some tests at some point...
	router := chi.NewRouter()

	// @MarkFix use other middleware - logger? recoverer?
	// @MarkFix CORS? Use middleware

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
	router.Get("/auth", handlers.Make(handlers.HandleAuthLogin(authObj)))
	router.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback(authObj)))
	router.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout(authObj)))
	router.Get("/auth/logout/callback", handlers.Make(handlers.HandleAuthLogoutCallback))

	// pages
	router.Get("/", handlers.Make(handlers.HandleHome(authObj)))
	router.Get("/user", handlers.Make(handlers.HandleUser(authObj)))
	// @MarkFix create user page

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	// log details about host / ports / @hotreload dev watch proxies
	publicPort := utils.Getenv("PUBLIC_PORT")
	slog.Info("Starting http server", "URL", fmt.Sprintf("%s:%s",utils.Getenv("PUBLIC_HOST"), publicPort))
	if os.Getenv("TEMPL_WATCH_PROXY_URL") == utils.GetPublicURL() {
		slog.Info("Auth configured for watch proxy", "templWatchProxyUrl", os.Getenv("TEMPL_WATCH_PROXY_URL"))
		if utils.IsProd {
			log.Fatal("Why is TEMPL_WATCH_PROXY_URL env variable set in prod?")
		}
	}

	// @MarkFix use ListenAndServeTLS
	err = http.ListenAndServe(":"+publicPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	log.Println("http server stopped")
}
