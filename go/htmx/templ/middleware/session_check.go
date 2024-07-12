package middleware

import (
	"database/sql"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/store"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/oauth2"
)

// Checks whether the user has a session. If they do, validate it and/or refresh token
func SessionCheck(authObj *auth.Authenticator, db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookieSession := store.GetSession(r)
			
			sessionID := cookieSession.Values["session_id"]
			userID := cookieSession.Values["user_id"]

			if sessionID == nil || userID == nil {
				DeleteSessionCookies(cookieSession, w, r) // ensure both are nil // @MarkFix could avoid this for performance
				next.ServeHTTP(w, r)
				return
			}

			var dbSession *model.Session = nil
			dbSession = database.GetSession(db, sessionID.(string), userID.(string))

			if dbSession == nil {
				DeleteSessionCookies(cookieSession, w, r)
				next.ServeHTTP(w, r)
				return
			}

			restoredToken := &oauth2.Token{
				AccessToken:  dbSession.AccessToken,
				RefreshToken: dbSession.RefreshToken,
				Expiry:       time.Unix(int64(dbSession.Expiry), 0),
				TokenType:    dbSession.TokenType,
			}

			tokenSource := authObj.TokenSource(r.Context(), restoredToken)
			newToken, err := tokenSource.Token()
			if err != nil {
				// something was wrong with the token and/or it failed to refresh
				/* @MarkFix to get here:
				   - lower "Access Token Lifespan": http://0.0.0.0:8080/admin/master/console/#/templ-realm/clients/e49746db-625e-49cd-91af-afbf8833bf95/advanced
				   - delete session from keycloak:  http://0.0.0.0:8080/admin/master/console/#/templ-realm/users/294e636d-88d7-4a62-88c2-4e9c6b0616e0/sessions
				   - wait for token to expire
				   - visit home
				*/
				DeleteSession(db, dbSession.ID, cookieSession, w, r)
				next.ServeHTTP(w, r)
				return
				// @MarkFix also redirect here?
			}

			// if the token has been refreshed, verify the IDToken 
			if newToken.AccessToken != restoredToken.AccessToken {
				_, err = authObj.VerifyIDToken(r.Context(), newToken) // @MarkFix I need verify the nonce? According to func's code comment I do
				if err != nil {
					// @MarkFix to get here, token needs to be refreshed but IDToken not valid?
					DeleteSession(db, dbSession.ID, cookieSession, w, r)
					next.ServeHTTP(w, r)
					return
				}
				// @MarkFix need to update the user profile here? Shouldn't be possible for a different user to refresh the token?
				// @MarkFix maybe just assert old userID == newUserID
			}

			// if the token has been refreshed, update session
			if dbSession.AccessToken != newToken.AccessToken {
				// @MarkFix we could update the sessionID here, to be paranoid...
				database.UpdateSession(db, sessionID.(string), userID.(string), newToken.AccessToken,
					newToken.RefreshToken, newToken.Expiry.Unix(), newToken.TokenType)
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func DeleteSessionCookies(cookieSession *sessions.Session, w http.ResponseWriter, r *http.Request) {
	cookieSession.Values["session_id"] = nil
	cookieSession.Values["user_id"] = nil
	store.SaveSession(cookieSession, w, r)
}

func DeleteSession(db *sql.DB, dbSessionID string, cookieSession *sessions.Session, w http.ResponseWriter, r *http.Request) {
	database.DeleteSession(db, dbSessionID)

	DeleteSessionCookies(cookieSession, w, r)

	auth.RawIDToken = ""
	auth.Profile = nil
}
