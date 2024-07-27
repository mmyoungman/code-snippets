package handlers

import (
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleExamples() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		userUntyped := r.Context().Value(utils.UserCtxKey)
		if userUntyped != nil {
			user := userUntyped.(*model.User)
			firstName = user.FirstName
		}

		return pages.Examples(firstName).Render(r.Context(), w)
	}
}
