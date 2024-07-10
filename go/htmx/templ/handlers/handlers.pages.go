package handlers

import (
	"mmyoungman/templ/auth"
	"mmyoungman/templ/views/pages"
	"net/http"
)


func HandleHome(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		if auth.Profile["given_name"] != nil { // @MarkFix get from DB
			firstName = auth.Profile["given_name"].(string)
		}
		return pages.Home(firstName).Render(r.Context(), w)
	}
}

func HandleUser(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		if auth.Profile["given_name"] != nil {
			firstName = auth.Profile["given_name"].(string)
		}
		return pages.User(firstName).Render(r.Context(), w)
	}
}