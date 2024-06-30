package handlers

import (
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleIndex(writer http.ResponseWriter, request *http.Request) error {
	return pages.Index().Render(request.Context(), writer)
}

func HandleLogin(writer http.ResponseWriter, request *http.Request) error {
	return pages.Login().Render(request.Context(), writer)
}

func HandleSignUp(writer http.ResponseWriter, request *http.Request) error {
	return pages.SignUp().Render(request.Context(), writer)
}