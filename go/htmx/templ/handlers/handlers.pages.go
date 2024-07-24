package handlers

import (
	"database/sql"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleHome(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		userUntyped := r.Context().Value(utils.ReqUserCtxKey)
		if userUntyped != nil {
			user := userUntyped.(*model.User)
			firstName = user.FirstName
		}

		return pages.Home(firstName).Render(r.Context(), w)
	}
}

func HandleUser(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		var user *model.User = nil
		userUntyped := r.Context().Value(utils.ReqUserCtxKey)
		if userUntyped != nil {
			user = userUntyped.(*model.User)
			return pages.UserLoggedIn(user).Render(r.Context(), w)
		}

		return pages.UserLoggedOut().Render(r.Context(), w)
	}
}
