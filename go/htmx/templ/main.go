package main

import (
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/handlers"
	"mmyoungman/templ/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db := database.DBConnect()
	// NOTE: db.Close() called by sqliteStore

	router := chi.NewMux()

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
	sqliteStore := auth.Setup(db)
	defer sqliteStore.Close()
	router.Get("/auth", handlers.Make(handlers.HandleAuthLogin))
	router.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback))
	router.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout))

	// pages
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/sign-up", handlers.Make(handlers.HandleSignUp))

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	// @MarkFix CORS?

	listenPort := utils.Getenv("PUBLIC_PORT")
	slog.Info("Starting http server", "listenPort", listenPort)
	// @MarkFix use ListenAndServeTLS
	err := http.ListenAndServe(":"+listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
