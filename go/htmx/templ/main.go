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
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Didn't load env file", err)
	}

	db := database.DBConnect()
	db.Close()

	auth.Setup(db) // @MarkFix we're not defer closing sqlitestore stuff

	// Routes
	router := chi.NewMux()

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
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

	port := utils.Getenv("PUBLIC_PORT")
	slog.Info("Starting http server", "URL", utils.Getenv("PUBLIC_HOST") + ":" + port)
	if os.Getenv("TEMPL_WATCH_PROXY_URL") != "" {
		slog.Info("Auth configured for watch proxy", "templWatchProxyUrl", utils.Getenv("TEMPL_WATCH_PROXY_URL"))
	}
	// @MarkFix use ListenAndServeTLS
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
