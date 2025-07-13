package store

import (
	"log"
	"mmyoungman/templ/utils"
	"net/http"

	"github.com/gorilla/sessions"
)

const SessionCookieName = "session"

var store *sessions.CookieStore

func Setup() {
	newStore := sessions.NewCookieStore(
		[]byte(utils.Getenv("SESSION_SECRET")))
	//[]byte(utils.Getenv("SESSION_SECRET"))) // @MarkFix additional arg for encryption?
	// @MarkFix review all cookie options
	newStore.Options.Path = "/"
	newStore.Options.SameSite = http.SameSiteLaxMode
	newStore.Options.HttpOnly = true
	//newStore.Options.Secure = utils.IsProd // @MarkFix we should be using TLS on prod, right?

	store = newStore
}

func GetSession(r *http.Request, name string) *sessions.Session {
	session, err := store.Get(r, name)
	if err != nil {
		log.Fatal("Error in fetching cookie session ", err)
	}
	return session
}

func SaveSession(session *sessions.Session, w http.ResponseWriter, r *http.Request) {
	err := session.Save(r, w)
	if err != nil {
		log.Fatal("Failed to save cookie session ", err)
	}
}

func DeleteSession(cookieSession *sessions.Session, w http.ResponseWriter, r *http.Request) {
	cookieSession.Values["session_id"] = nil
	cookieSession.Values["state_login"] = nil
	cookieSession.Values["state_logout"] = nil
	cookieSession.Values["pkce_verifier"] = nil
	cookieSession.Values["referrer_path"] = nil
	cookieSession.Values["csrf_token"] = nil

	cookieSession.Options.MaxAge = -1

	SaveSession(cookieSession, w, r)
}
