package main

import (
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/handlers"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Get("/test", handlers.Make(handlers.HandleTest))

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("Starting http server", "listenAddr", listenAddr)

	err := http.ListenAndServe(listenAddr, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}

