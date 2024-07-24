package main

import (
	"embed"
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/handlers"
	"mmyoungman/templ/middleware"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

//go:embed database/migrations/*
var embedMigrations embed.FS

func main() {
	// include file and line in log messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// @MarkFix make the program print all logs to a file

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Didn't load env file", err)
	}

	db := database.Connect()
	defer db.Close()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("Failed to set goose dialect ", err)
	}
	if err := goose.Up(db, utils.Getenv("MIGRATIONS_PATH")); err != nil {
		log.Fatal("Failed to apply migrations ", err) // @MarkFix do we actually want to fail here?
	}

	authObj, err := auth.Setup()
	if err != nil {
		log.Fatal("Auth setup failed: ", err)
	}

	store.Setup()

	// @MarkFix build pipeline / deployment?
	// @MarkFix set up auto formatting
	// @MarkFix research static code analysis
	// @MarkFix I suppose I could write some tests at some point...
	router := chi.NewRouter()

	router.Use(middleware.ContentSecurityPolicy)

	// @MarkFix use other middleware - logger? recoverer?
	// @MarkFix rate limiting middleware?
	// @MarkFix caching middleware?
	// @MarkFix compression middleware?
	// @MarkFix CORS middleware - github.com/rs/cors
	// @MarkFix monitoring / analytics?
	// @MarkFix data backup / disaster recovery?

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
	router.Group(func(r chi.Router) {
		r.Use(middleware.SessionCheck(authObj, db))
		// we want to check whether user is already logged out in logout case
		r.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout(authObj, db)))
	})

	router.Group(func(r chi.Router) {
		r.Get("/auth", handlers.Make(handlers.HandleAuthLogin(authObj))) // @MarkFix do we want to check context user here?
		r.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback(authObj, db)))
		r.Get("/auth/logout/callback", handlers.Make(handlers.HandleAuthLogoutCallback(db)))
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.SessionCheck(authObj, db))

		// public pages (that have dynamic content depending on whether the user is logged in)
		r.Get("/", handlers.Make(handlers.HandleHome(authObj, db)))
		r.Get("/examples", handlers.Make(handlers.HandleExamples()))
		r.Get("/examples/click-button-load-partial", handlers.Make(handlers.HandleClickButtonLoadPartial()))
		r.Get("/examples/todo-list", handlers.Make(handlers.HandleToDoList(db)))

		// private pages (i.e. logged in users only)
		r.Get("/user", handlers.Make(handlers.HandleUser(authObj, db)))

		// partials
		r.Get("/test", handlers.Make(handlers.HandleTest))
		r.Get("/todo-item-list", handlers.Make(handlers.HandleToDoListItems(db)))
		r.Get("/todo-add-item-form", handlers.Make(handlers.HandleToDoAddForm()))
		r.Post("/todo-add-form-submit", handlers.Make(handlers.HandleToDoAddFormSubmit(db)))
		r.Get("/todo-update-item-form", handlers.Make(handlers.HandleToDoUpdateForm(db)))
		r.Put("/todo-update-form-submit", handlers.Make(handlers.HandleToDoUpdateFormSubmit(db)))
		r.Get("/todo-form-cancel", handlers.Make(handlers.HandleToDoFormCancel()))
	})

	// log details about host / ports / @hotreload dev watch proxies
	publicPort := utils.Getenv("PUBLIC_PORT")
	slog.Info("Starting http server", "URL", fmt.Sprintf("%s:%s", utils.Getenv("PUBLIC_HOST"), publicPort))
	if os.Getenv("PROXY_URL") == utils.GetPublicURL() {
		slog.Info("Auth configured for watch proxy", "proxyUrl", os.Getenv("PROXY_URL"))
		if utils.IsProd {
			log.Fatal("Why is PROXY_URL env variable set in prod?")
		}
	}

	if !utils.IsProd && (os.Getenv("TLS_CRT") == "" || os.Getenv("TLS_KEY") == "") {
		err = http.ListenAndServe(
			":"+publicPort,
			router)
		if err != nil {
			log.Fatal("ListenAndServe error: ", err)
		}
	} else {
		err = http.ListenAndServeTLS(
			":"+publicPort,
			utils.Getenv("TLS_CRT"),
			utils.Getenv("TLS_KEY"),
			router)
		if err != nil {
			log.Fatal("ListenAndServeTLS error: ", err)
		}
	}
	log.Println("Server stopped")
}
