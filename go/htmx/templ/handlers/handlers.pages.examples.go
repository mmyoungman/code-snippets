package handlers

import (
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/layouts"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleExamples() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := layouts.BaseArgs{
			Nonce: utils.GetContextCspNonce(r),
			CsrfToken: utils.GetContextCSRFToken(r),
		}
		user := utils.GetContextUser(r)
		if user != nil {
			baseArgs.Username = user.FirstName
		}

		return pages.Examples(baseArgs).Render(r.Context(), w)
	}
}
