package middleware

import (
	"log"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"

	"github.com/google/uuid"
)

func CSRFToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookieSession := store.GetSession(r, store.SessionCookieName)
		csrfToken := cookieSession.Values["csrf_token"]
		if csrfToken == nil {
			csrfToken = uuid.New().String()
			// @MarkFix this cookie currently never changes - it probably should for security
			// Removed it if the csrf token is used
			// or it could be saved in a separate cookie with a TTL of like 10 mins
			// or both
			cookieSession.Values["csrf_token"] = csrfToken
			store.SaveSession(cookieSession, w, r)
		}

		csrfTokenStr, csrfTokenIsString := csrfToken.(string)
		if !csrfTokenIsString {
			log.Fatal("CSRF token could not be cast into a string")
		}

		utils.SetContextValue(r, utils.CsrfTokenCtxKey, csrfTokenStr)
		next.ServeHTTP(w, r)
	})
}
