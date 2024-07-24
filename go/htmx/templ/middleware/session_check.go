package middleware

import (
	"context"
	"database/sql"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// Checks whether the user has a session. If they do, validate it and/or refresh token
func SessionCheck(authObj *auth.Authenticator, db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookieSession := store.GetSession(r)

			sessionIDUntyped := cookieSession.Values["session_id"]

			if sessionIDUntyped == nil {
				store.DeleteSession(cookieSession, w, r) // ensure both are nil // @MarkFix could avoid this for performance
				next.ServeHTTP(w, r)
				return
			}

			var dbSession *model.Session = nil
			dbSession = database.GetSession(db, sessionIDUntyped.(string))

			if dbSession == nil {
				store.DeleteSession(cookieSession, w, r)
				next.ServeHTTP(w, r)
				return
			}

			restoredToken := &oauth2.Token{
				AccessToken:  dbSession.AccessToken,
				RefreshToken: dbSession.RefreshToken,
				Expiry:       time.Unix(int64(dbSession.Expiry), 0),
				TokenType:    dbSession.TokenType,
			}
			//var tokenExtraMap = make(map[string]string) // @MarkFix doesn't work?
			//tokenExtraMap["id_token"] = auth.RawIDToken
			//restoredToken.WithExtra(tokenExtraMap)

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

				database.DeleteSession(db, dbSession.ID)
				store.DeleteSession(cookieSession, w, r)

				// @MarkFix don't delete user?

				next.ServeHTTP(w, r)
				return
			}

			// if the token has been refreshed, verify the IDToken
			if newToken.AccessToken != restoredToken.AccessToken {
				_, err = authObj.VerifyIDToken(r.Context(), newToken) // @MarkFix I need verify the nonce? According to func's code comment I do
				if err != nil {
					// @MarkFix to get here, token needs to be refreshed but IDToken not valid?

					database.DeleteSession(db, dbSession.ID)
					store.DeleteSession(cookieSession, w, r)

					// @MarkFix don't delete user?

					next.ServeHTTP(w, r)
					return
				}
				// @MarkFix need to update the user profile here? Shouldn't be possible for a different user to refresh the token?
				// @MarkFix maybe just assert old userID == newUserID
			}

			// if the token has been refreshed, update session
			if dbSession.AccessToken != newToken.AccessToken {
				// @MarkFix we could update the sessionID here, to be paranoid...
				database.UpdateSession(db, dbSession.ID, dbSession.UserID, newToken.AccessToken,
					newToken.RefreshToken, newToken.Expiry.Unix(), newToken.TokenType)
			}

			// user is logged in, so set user on context for access in handlers
			loggedInUser := database.GetUser(db, dbSession.UserID) // @MarkFix _should_ always return a user - add db relationship to ensure or double check here
			ctx := r.Context()
			newCtx := context.WithValue(ctx, utils.ReqUserCtxKey, loggedInUser)
			*r = *r.WithContext(newCtx)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
