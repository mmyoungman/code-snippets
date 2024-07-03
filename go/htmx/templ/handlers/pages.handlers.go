package handlers

import (
	"mmyoungman/templ/views/pages"
	"net/http"

	"github.com/markbates/goth/gothic"
)

func HandleHome(w http.ResponseWriter, r *http.Request) error {
	username, err := gothic.GetFromSession("username", r)
	if err != nil {
		return pages.Home("").Render(r.Context(), w)
	}
	return pages.Home(username).Render(r.Context(), w)
}

func HandleLogin(w http.ResponseWriter, r *http.Request) error {
	return pages.Login().Render(r.Context(), w)
}

func HandleSignUp(w http.ResponseWriter, r *http.Request) error {
	return pages.SignUp().Render(r.Context(), w)
}