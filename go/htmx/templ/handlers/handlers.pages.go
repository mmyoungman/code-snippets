package handlers

import (
	"database/sql"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/layouts"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleHome() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := layouts.BaseArgs{
			Nonce: utils.GetContextCspNonce(r),
			CsrfToken: utils.GetContextCSRFToken(r),
		}
		user := utils.GetContextUser(r)
		if user != nil {
			baseArgs.Username = user.FirstName
		}

		return pages.Home(baseArgs).Render(r.Context(), w)
	}
}

func HandleUser(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := layouts.BaseArgs{
			Nonce: utils.GetContextCspNonce(r),
			CsrfToken: utils.GetContextCSRFToken(r),
		}

		user := utils.GetContextUser(r)

		if user == nil {
			return pages.UserLoggedOut(baseArgs).Render(r.Context(), w)
		}

		baseArgs.Username = user.FirstName
		return pages.UserLoggedIn(user, baseArgs).Render(r.Context(), w)
	}
}
