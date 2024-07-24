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

func HandleToDoAddForm() HTTPHandler {
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
		return pages.UpdatePageAfterAddFormSubmit(newItem).Render(r.Context(), w)
	}
}

func HandleToDoUpdateForm(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.URL.Query().Get("id") // @MarkFix validation

		item := database.GetToDoItem(db, id)

		// @MarkFix can visit {URL}/test directly in a browser
		return pages.UpdateItemForm(item).Render(r.Context(), w)
	}
}

func HandleToDoUpdateFormSubmit(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		r.ParseForm()

		// @MarkFix validate form values

		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")

		newItem := database.UpdateToDoItem(db, &model.ToDoItem{
			ID:          id,
			Name:        name,
			Description: description,
		})

		// @MarkFix can visit {URL}/test directly in a browser
		return pages.UpdatePageAfterUpdateFormSubmit(newItem).Render(r.Context(), w)
	}
}

func HandleToDoFormCancel() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return pages.AddItemButton().Render(r.Context(), w)
	}
}
