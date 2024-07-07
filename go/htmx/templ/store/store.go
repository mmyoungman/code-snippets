package store

import (
	"mmyoungman/templ/utils"

	"github.com/gorilla/sessions"
)

var store *sessions.CookieStore

func Setup() {
	newStore := sessions.NewCookieStore(
		[]byte(utils.Getenv("SESSION_SECRET")))
	//store.Options.HttpOnly = true
	store = newStore

	utils.UNUSED(store)
}