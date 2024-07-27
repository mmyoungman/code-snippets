package handlers

import (
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleExamples() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		user := utils.GetContextUser(r)
		if user != nil {
			firstName = user.FirstName
		}

		return pages.Examples(firstName, utils.GetContextCspNonce(r)).Render(r.Context(), w)
	}
}
