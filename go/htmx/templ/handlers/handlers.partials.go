package handlers

import (
	"mmyoungman/templ/views/partials"
	"net/http"
)

func HandleTest(writer http.ResponseWriter, request *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return partials.Test().Render(request.Context(), writer)
}
