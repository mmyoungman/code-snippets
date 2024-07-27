package handlers

import (
	"database/sql"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleHome() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		user := utils.GetContextUser(r)
		if user != nil {
			firstName = user.FirstName
		}

		return pages.Home(firstName, utils.GetContextCspNonce(r)).Render(r.Context(), w)
	}
}

func HandleUser(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		user := utils.GetContextUser(r)
		if user != nil {
			return pages.UserLoggedIn(user, utils.GetContextCspNonce(r)).Render(r.Context(), w)
		}

		return pages.UserLoggedOut(utils.GetContextCspNonce(r)).Render(r.Context(), w)
	}
}
