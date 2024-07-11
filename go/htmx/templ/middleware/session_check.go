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

	"golang.org/x/oauth2"
)

// Checks whether the user has a session. If they do, validate it and/or refresh token
func SessionCheck(authObj *auth.Authenticator, db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			// @MarkFix If someone changed Token.Expiry to far in the future, we'd currently
			// never revalidate the user and they'd stay logged in permanently? So maybe
			// VerifyIDToken every time?
			// @MarkFix all this logic needs double checking once using a db to store session data
			cookieSession := store.GetSession(r)
			
			sessionID := cookieSession.Values["session_id"]
			userID := cookieSession.Values["user_id"]

			var dbSession *model.Sessions = nil
			if sessionID != nil && userID != nil {
				dbSession = database.GetSession(db, sessionID.(string), userID.(string))
				if dbSession != nil {
					fmt.Println("Expiry: ", time.Unix(int64(dbSession.Expiry), 0).String())
				}
			}

			if dbSession != nil && int64(dbSession.Expiry) > time.Now().Unix() {
				restoredToken := &oauth2.Token{
					AccessToken:  dbSession.AccessToken,
					RefreshToken: dbSession.RefreshToken,
					Expiry:       time.Unix(int64(dbSession.Expiry), 0),
					TokenType:    dbSession.TokenType,

				}

				tokenSource := authObj.TokenSource(r.Context(), restoredToken)

				newToken, err := tokenSource.Token()
				if err != nil {
					log.Fatal("Refresh token failed or something: ", err)
					/* @MarkFix to get here:
					   - lower "Access Token Lifespan": http://0.0.0.0:8080/admin/master/console/#/templ-realm/clients/e49746db-625e-49cd-91af-afbf8833bf95/advanced
					   - delete session from keycloak:  http://0.0.0.0:8080/admin/master/console/#/templ-realm/users/294e636d-88d7-4a62-88c2-4e9c6b0616e0/sessions
					   - wait for token to expire
					   - visit home
					   The token cannot refresh - so what should we do here?
					   - Set auth.Token to null and/or delete all session data? Maybe even delete user data?
					   - Redirect to login and then back to this page?
					   - Maybe both?
					*/
				}
				_, err = authObj.VerifyIDToken(r.Context(), newToken) // @MarkFix do we bother with this? - and I need verify the nonce?
				if err != nil {
					log.Fatal("Failed to verify IDToken ", err)
					// @MarkFix Is this right?
					database.DeleteSession(db, dbSession.ID)

					cookieSession.Values["session_id"] = nil
					cookieSession.Values["user_id"] = nil
					if err := cookieSession.Save(r, w); err != nil {
						log.Fatal("Failed to save cookieSession after finding bad accesstoken", err)
					}

					auth.RawIDToken = ""
					auth.Profile = nil

					next.ServeHTTP(w, r)
					return
				}

				if dbSession.AccessToken == newToken.AccessToken {
					log.Fatal("Something has gone wrong?")
				}

				database.UpdateSession(db, sessionID.(string), userID.(string), newToken.AccessToken,
					newToken.RefreshToken, newToken.Expiry.Unix(), newToken.TokenType)
			}

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
