package handlers

import (
	"database/sql"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleClickButtonLoadPartial() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		userUntyped := r.Context().Value(utils.ReqUserCtxKey)
		if userUntyped != nil {
			user := userUntyped.(*model.User)
			firstName = user.FirstName
		}

		return pages.ExamplesClickButtonLoadPartial(firstName).Render(r.Context(), w)
	}
}

func HandleToDoList(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		firstName := ""
		userUntyped := r.Context().Value(utils.ReqUserCtxKey)
		if userUntyped != nil {
			user := userUntyped.(*model.User)
			firstName = user.FirstName
		}

		toDoItems := database.ListToDoItems(db)

		return pages.ExamplesToDoList(firstName, toDoItems).Render(r.Context(), w)
	}
}
