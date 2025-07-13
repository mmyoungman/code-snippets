package handlers

import (
	"database/sql"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/layouts"
	"mmyoungman/templ/views/pages"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func HandleToDoList(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := layouts.BaseArgs{
			Nonce: utils.GetContextCspNonce(r),
			CsrfToken: utils.GetContextCSRFToken(r),
		}
		user := utils.GetContextUser(r)
		if user != nil {
			baseArgs.Username = user.FirstName
		}

		toDoItems := database.ListToDoItems(db)

		return pages.ExamplesToDoList(baseArgs, toDoItems).Render(r.Context(), w)
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
		id = strings.TrimPrefix(id, "item-")

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

func HandleToDoDelete(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.URL.Query().Get("id") // @MarkFix validation
		id = strings.TrimPrefix(id, "item-")

		database.DeleteToDoItem(db, id)

		return pages.DeleteToDoItem().Render(r.Context(), w)
	}
}

func HandleToDoFormCancel() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return pages.DefaultControls().Render(r.Context(), w)
	}
}
