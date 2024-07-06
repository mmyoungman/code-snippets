package handlers

import (
	"mmyoungman/templ/views/pages"
	"mmyoungman/templ/auth"
	"net/http"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	firstName := ""
	if auth.Profile["given_name"] != nil {
		firstName = auth.Profile["given_name"].(string)
	}
	//return pages.Home(username).Render(r.Context(), w)
	return pages.Home(firstName).Render(r.Context(), w)
}