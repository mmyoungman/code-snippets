package handlers

import (
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleExamples() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := utils.GenerateBaseArgs(r)
		return pages.Examples(baseArgs).Render(r.Context(), w)
	}
}
