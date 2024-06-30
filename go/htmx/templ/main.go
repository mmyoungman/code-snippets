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

	router.Handle("/*", public())

	router.Get("/", handlers.Make(handlers.HandleIndex))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/sign-up", handlers.Make(handlers.HandleSignUp))

	listenPort := utils.Getenv("LISTEN_PORT")

	slog.Info("Starting http server", "listenPort", listenPort)

	err := http.ListenAndServe(":" + listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}