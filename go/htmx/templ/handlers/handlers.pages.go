package handlers

import (
	"database/sql"
	"log"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/store"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleHome(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""

		// @MarkFix Need to check dbSession? We've already run SessionCheck...
		session := store.GetSession(r)
		userID := session.Values["user_id"]
		if userID != nil {
			user := database.GetUser(db, userID.(string))
			if user == nil {
				log.Fatal("Wrong") // @MarkFix handle this
			}
			firstName = user.FirstName
		}

		return pages.Home(firstName).Render(r.Context(), w)
	}
}

func HandleUser(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""

		// @MarkFix Need to check dbSession? We've already run SessionCheck...
		session := store.GetSession(r)
		userID := session.Values["user_id"]
		if userID != nil {
			user := database.GetUser(db, userID.(string))
			if user == nil {
				log.Fatal("Wrong") // @MarkFix handle this
			}
			firstName = user.FirstName
		}

		return pages.User(firstName).Render(r.Context(), w)
	}
}
