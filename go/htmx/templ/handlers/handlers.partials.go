package handlers

import (
	"database/sql"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/views/partials"
	"net/http"

	"github.com/google/uuid"
)

func HandleTest(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return partials.Test().Render(r.Context(), w)
}

func HandleToDoListAdd(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		r.ParseForm()

		// @MarkFix validate form values

		name := r.FormValue("name")
		description := r.FormValue("description")

		database.InsertToDoItem(db, &model.ToDoItem{ 
			ID: uuid.NewString(), 
			Name: name, 
			Description: description,
		})

		// @MarkFix can visit {URL}/test directly in a browser
		return partials.ToDoItemRow(name, description).Render(r.Context(), w)
	}
}