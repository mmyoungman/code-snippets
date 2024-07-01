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

	// @MarkFix auth routes
	store := sessions.NewCookieStore([]byte("sessionkey"))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false

	gothic.Store = store

	openidConnect, err := openidConnect.New(
		utils.Getenv("KEYCLOAK_CLIENTID"),
		utils.Getenv("KEYCLOAK_OIDC_SECRET"),
		"http://localhost:3000/auth/callback?provider=openid-connect",
		utils.Getenv("KEYCLOAK_DISCOVERY_URL"))
	if err != nil {
		log.Fatal("Is keycloak started? Error:\n", err)
	}
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	// auth
	router.Get("/auth/callback", handlers.Make(handlers.HandleAuthCallback))
	router.Get("/auth/logout", handlers.Make(handlers.HandleAuthLogout))
	router.Get("/auth", handlers.Make(handlers.HandleAuthLogin))

	// pages
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/sign-up", handlers.Make(handlers.HandleSignUp))

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	listenPort := utils.Getenv("LISTEN_PORT")
	slog.Info("Starting http server", "listenPort", listenPort)
	// @MarkFix use ListenAndServeTLS
	err = http.ListenAndServe(":"+listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
