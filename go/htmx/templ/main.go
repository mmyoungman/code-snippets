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

	router.Handle("/*", public())

	router.Get("/test", handlers.Make(handlers.HandleTest))

	listenPort := os.Getenv("LISTEN_PORT")
	//if listenPort == "" {
	//	log.Fatal("LISTEN_PORT not defined in .env file")
	//}

	slog.Info("Starting http server", "listenPort", listenPort)

	err := http.ListenAndServe(":" + listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}

