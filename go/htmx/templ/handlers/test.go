package handlers

import (
	"mmyoungman/templ/views/test"
	"net/http"
)

func HandleTest(writer http.ResponseWriter, request *http.Request) error {
	return test.Index().Render(request.Context(), writer)
}
