package handlers

import (
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleClickButtonLoadPartial() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		user := utils.GetContextUser(r)
		if user != nil {
			firstName = user.FirstName
		}

		return pages.ExamplesClickButtonLoadPartial(firstName, utils.GetContextCspNonce(r)).Render(r.Context(), w)
	}
}

func HandleTest(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.Test().Render(r.Context(), w)
}
