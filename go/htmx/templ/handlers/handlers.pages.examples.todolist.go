package handlers

import (
	"context"
	"database/sql"
	"mmyoungman/templ/database/sqlc_gen"
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func HandleToDoList(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := utils.GenerateBaseArgs(r)

		queries := database.New(db)
		toDoItems, err := queries.ListToDoItems(context.Background())

		utils.UNUSED(err) // @MarkFix handle err

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

		queries := database.New(db)
		newItem, err := queries.InsertToDoItem(context.Background(), database.InsertToDoItemParams{
			ID:          uuid.NewString(),
			Name:        name,
			Description: description,
		})

		utils.UNUSED(err) // @MarkFix handle err

		// @MarkFix can visit {URL}/test directly in a browser
		return pages.UpdatePageAfterAddFormSubmit(newItem).Render(r.Context(), w)
	}
}

func HandleToDoUpdateForm(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.URL.Query().Get("id") // @MarkFix validation
		id = strings.TrimPrefix(id, "item-")

		queries := database.New(db)
		item, err := queries.GetToDoItem(context.Background(), id)

		utils.UNUSED(err) // @MarkFix handle err

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

		queries := database.New(db)
		newItem, err := queries.UpdateToDoItem(context.Background(), database.UpdateToDoItemParams{
			ID:          id,
			Name:        name,
			Description: description,
		})

		utils.UNUSED(err) // @MarkFix handle err?

		// @MarkFix can visit {URL}/test directly in a browser
		return pages.UpdatePageAfterUpdateFormSubmit(newItem).Render(r.Context(), w)
	}
}

func HandleToDoDelete(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		id := r.URL.Query().Get("id") // @MarkFix validation
		id = strings.TrimPrefix(id, "item-")

		queries := database.New(db)
		queries.DeleteToDoItem(context.Background(), id)

		return pages.DeleteToDoItem().Render(r.Context(), w)
	}
}

func HandleToDoFormCancel() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		return pages.DefaultControls().Render(r.Context(), w)
	}
}
