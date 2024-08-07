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
	"mmyoungman/templ/structs"
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

	// @MarkFix create a sessionCtx object? (or use r.Context()?)

	serviceCtx := &structs.ServiceCtx{}

	serviceCtx.Db = database.Connect()
	defer serviceCtx.Db.Close()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatal("Failed to set goose dialect ", err)
	}
	if err := goose.Up(serviceCtx.Db, utils.Getenv("MIGRATIONS_PATH")); err != nil {
		log.Fatal("Failed to apply migrations ", err) // @MarkFix do we actually want to fail here?
	}

	var err error
	serviceCtx.Auth, err = auth.Setup()
	if err != nil {
		log.Fatal("Auth setup failed: ", err)
	}

	store.Setup()

	// @MarkFix build pipeline / deployment?
	// @MarkFix set up auto formatting
	// @MarkFix research static code analysis
	// @MarkFix I suppose I could write some tests at some point...
	router := chi.NewRouter()

	router.Use(middleware.ContentSecurityPolicy) // @MarkFix console errors due to this?

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
		r.Use(middleware.SessionCheck(serviceCtx))
		// we want to check whether user is already logged out in logout case
		r.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout(serviceCtx)))
	})

	router.Group(func(r chi.Router) {
		r.Get("/auth", handlers.Make(handlers.HandleAuthLogin(serviceCtx.Auth)))
		r.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback(serviceCtx)))
		r.Get("/auth/logout/callback", handlers.Make(handlers.HandleAuthLogoutCallback(serviceCtx.Db)))
	})

	router.Group(func(r chi.Router) {
		r.Use(middleware.SessionCheck(serviceCtx))

		// public pages (that have dynamic content depending on whether the user is logged in)
		r.Get("/", handlers.Make(handlers.HandleHome()))
		r.Get("/examples", handlers.Make(handlers.HandleExamples()))
		r.Get("/examples/click-button-load-partial", handlers.Make(handlers.HandleClickButtonLoadPartial()))
		r.Get("/examples/todo", handlers.Make(handlers.HandleToDoList(serviceCtx.Db)))

		// private pages (i.e. logged in users only)
		r.Get("/user", handlers.Make(handlers.HandleUser(serviceCtx.Auth, serviceCtx.Db)))

		// partials
		r.Get("/partial/test1", handlers.Make(handlers.HandleTest1))
		r.Get("/partial/test2", handlers.Make(handlers.HandleTest2))
		r.Get("/partial/place-button-in-target1", handlers.Make(handlers.HandlePlaceButtonInTarget1))
		r.Get("/partial/place-button-in-target2", handlers.Make(handlers.HandlePlaceButtonInTarget2))

		r.Get("/todo/add-form", handlers.Make(handlers.HandleToDoAddForm()))
		r.Post("/todo", handlers.Make(handlers.HandleToDoAddFormSubmit(serviceCtx.Db)))
		r.Get("/todo/update-form", handlers.Make(handlers.HandleToDoUpdateForm(serviceCtx.Db)))
		r.Put("/todo", handlers.Make(handlers.HandleToDoUpdateFormSubmit(serviceCtx.Db)))
		r.Delete("/todo", handlers.Make(handlers.HandleToDoDelete(serviceCtx.Db)))
		r.Get("/todo/form-cancel", handlers.Make(handlers.HandleToDoFormCancel()))
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
