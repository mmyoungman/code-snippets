package store

import (
	"log"
	"mmyoungman/templ/utils"
	"net/http"

	"github.com/gorilla/sessions"
)

const Name = "session"
var store *sessions.CookieStore

func Setup() {
	newStore := sessions.NewCookieStore(
		[]byte(utils.Getenv("SESSION_SECRET")))
		//[]byte(utils.Getenv("SESSION_SECRET"))) // @MarkFix additional arg for encryption?
	newStore.Options.Path = "/"
	//newStore.Options.HttpOnly = true
	//newStore.Options.Secure = utils.IsProd // we should be using TLS on prod, right?

	store = newStore
}

func GetSession(r *http.Request) *sessions.Session {
	session, err := store.Get(r, "session")
	if err != nil {
		log.Fatal("error in fetching session - should always return a session?", err)
	}
	return session
}