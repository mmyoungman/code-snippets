package handlers

import (
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleClickButtonLoadPartial() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		userUntyped := r.Context().Value(utils.UserCtxKey)
		if userUntyped != nil {
			user := userUntyped.(*model.User)
			firstName = user.FirstName
		}

		return pages.ExamplesClickButtonLoadPartial(firstName).Render(r.Context(), w)
	}
}

func HandleTest(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.Test().Render(r.Context(), w)
}
