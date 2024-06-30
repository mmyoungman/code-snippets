package handlers

import (
	"mmyoungman/templ/views/partials"
	"net/http"
)

func HandleTest(writer http.ResponseWriter, request *http.Request) error {
	return partials.Test().Render(request.Context(), writer)
}