package main

import (
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/database"
	"mmyoungman/templ/handlers"
	"mmyoungman/templ/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/openidConnect"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db := database.DBConnect()
	defer db.Close()

	router := chi.NewMux()

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// auth
	fileStore := sessions.NewFilesystemStore("./tmp/", []byte(utils.Getenv("SESSION_SECRET")))
	fileStore.MaxLength(8192)
	gothic.Store = fileStore

	openidConnect, err := openidConnect.New(
		utils.Getenv("KEYCLOAK_CLIENT_ID"),
		utils.Getenv("KEYCLOAK_CLIENT_SECRET"),
		utils.Getenv("PUBLIC_URL") + "/auth/callback?provider=openid-connect",
		utils.Getenv("KEYCLOAK_DISCOVERY_URL"))
	if err != nil {
		log.Fatal("Error creating openidConnect provider. Is keycloak started? Error:\n", err)
	}
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

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
	err = http.ListenAndServe(":"+listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
