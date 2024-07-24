package handlers

import (
	"mmyoungman/templ/views/partials"
	"net/http"
)

func HandleTest(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return partials.Test().Render(r.Context(), w)
}

func HandleToDoListAdd(w http.ResponseWriter, r *http.Request) error {
	r.ParseForm()

	// @MarkFix validate form values

	name := r.FormValue("name")
	description := r.FormValue("description")

	// @MarkFix add to database

	// @MarkFix can visit {URL}/test directly in a browser
	return partials.ToDoItemRow(name, description).Render(r.Context(), w)
}