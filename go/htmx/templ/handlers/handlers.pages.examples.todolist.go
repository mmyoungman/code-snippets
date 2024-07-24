package handlers

import (
	"database/sql"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"

	"github.com/google/uuid"
)

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

func HandleToDoListItems(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		toDoItems := database.ListToDoItems(db)

		return pages.ToDoItemList(toDoItems).Render(r.Context(), w)
	}
}

func HandleToDoListAddForm() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// @MarkFix can visit {URL}/test directly in a browser
		return pages.AddItemForm().Render(r.Context(), w)
	}
}

func HandleToDoAddFormSubmit(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		r.ParseForm()

		// @MarkFix validate form values

		name := r.FormValue("name")
		description := r.FormValue("description")

		newItem := database.InsertToDoItem(db, &model.ToDoItem{
			ID:          uuid.NewString(),
			Name:        name,
			Description: description,
		})

		// @MarkFix can visit {URL}/test directly in a browser
		return pages.UpdatePageAfterFormSubmit(newItem).Render(r.Context(), w)
	}
}

func HandleToDoFormCancel() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return pages.AddItemButton().Render(r.Context(), w)
	}
}
