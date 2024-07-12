package middleware

import (
	"database/sql"
	"fmt"
	"log"
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
			// @MarkFix all this logic needs double checking once using a db to store session data
			cookieSession := store.GetSession(r)
			
			sessionID := cookieSession.Values["session_id"]
			userID := cookieSession.Values["user_id"]

			var dbSession *model.Sessions = nil
			if sessionID != nil && userID != nil {
				dbSession = database.GetSession(db, sessionID.(string), userID.(string))
				if dbSession != nil {
					fmt.Println("SessionCheck Sess Found in DB, Expiry: ", time.Unix(int64(dbSession.Expiry), 0))
				}
			}

			if dbSession != nil {
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
					DeleteSession(db, dbSession.ID, cookieSession, r, w)
					next.ServeHTTP(w, r)
					return
					// @MarkFix also redirect here?
				}

				// if the token has been refreshed, verify the IDToken 
				if newToken.AccessToken != restoredToken.AccessToken {
					_, err = authObj.VerifyIDToken(r.Context(), newToken) // @MarkFix I need verify the nonce? According to func's code comment I do
					if err != nil {
						// @MarkFix to get here, token needs to be refreshed by IDToken not valid?
						DeleteSession(db, dbSession.ID, cookieSession, r, w)
						next.ServeHTTP(w, r)
						return
					}
					// @MarkFix need to reset the userID here? Shouldn't be possible for a different user to refresh the token?
					// @MarkFix maybe just assert old userID == newUserID
				}

				if dbSession.AccessToken != newToken.AccessToken {
					database.UpdateSession(db, sessionID.(string), userID.(string), newToken.AccessToken,
						newToken.RefreshToken, newToken.Expiry.Unix(), newToken.TokenType)
				}
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func DeleteSession(db *sql.DB, dbSessionID string, cookieSession *sessions.Session, r *http.Request, w http.ResponseWriter) {
	database.DeleteSession(db, dbSessionID)

	cookieSession.Values["session_id"] = nil
	cookieSession.Values["user_id"] = nil
	if err := cookieSession.Save(r, w); err != nil {
		log.Fatal("Failed to save cookieSession after finding bad accesstoken", err)
	}

	auth.RawIDToken = ""
	auth.Profile = nil
}
