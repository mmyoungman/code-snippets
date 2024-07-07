package handlers

import (
	"mmyoungman/templ/auth"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix if has sessions, and expired > Now, then use refresh token, update session info
	// using authObj.Client()?
	// And this should probably be done in middleware eventually

	firstName := ""
	if auth.Profile["given_name"] != nil {
		firstName = auth.Profile["given_name"].(string)
	}
	return pages.Home(firstName).Render(r.Context(), w)
}