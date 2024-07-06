package store

import (
	"mmyoungman/templ/utils"

	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func Setup() {
	store := sessions.NewCookieStore(
		[]byte(utils.Getenv("SESSION_SECRET")))
	//store.Options.HttpOnly = true
	Store = store
}