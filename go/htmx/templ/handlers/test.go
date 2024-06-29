package handlers

import (
	"mmyoungman/templ/views/test"
	"net/http"
)

func HandleTest(w http.ResponseWriter, r *http.Request) error {
	return Render(w, r, test.Index())
}
