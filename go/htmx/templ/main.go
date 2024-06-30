package main

import (
	"fmt"
	"log"
	"log/slog"
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

	router := chi.NewMux()

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// pages
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/sign-up", handlers.Make(handlers.HandleSignUp))

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	listenPort := utils.Getenv("LISTEN_PORT")
	slog.Info("Starting http server", "listenPort", listenPort)
	err := http.ListenAndServe(":" + listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}