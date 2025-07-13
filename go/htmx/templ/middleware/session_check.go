package middleware

import (
	"context"
	"database/sql"
	"log"
	"mmyoungman/templ/database/sqlc_gen"
	"mmyoungman/templ/store"
	"mmyoungman/templ/structs"
	"mmyoungman/templ/utils"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

// Checks whether the user has a session. If they do, validate it and/or refresh token
func SessionCheck(serviceCtx *structs.ServiceCtx) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			cookieSession := store.GetSession(r, store.SessionCookieName)

			sessionIDUntyped := cookieSession.Values["session_id"]

			if sessionIDUntyped == nil {
				store.DeleteSession(cookieSession, w, r) // ensure both are nil // @MarkFix could avoid this for performance
				utils.SetContextValue(r, utils.UserCtxKey, nil)

				next.ServeHTTP(w, r)
				return
			}

			queries := database.New(serviceCtx.Db)
			dbSession, err := queries.GetSession(context.Background(), sessionIDUntyped.(string))

			if err != nil {
				if err == sql.ErrNoRows {
					store.DeleteSession(cookieSession, w, r)
					utils.SetContextValue(r, utils.UserCtxKey, nil)

					next.ServeHTTP(w, r)
					return
				} else {
					log.Fatal("Error fetching sessions")
				}
			}

			restoredToken := &oauth2.Token{
				AccessToken:  dbSession.Accesstoken,
				RefreshToken: dbSession.Refreshtoken,
				Expiry:       time.Unix(int64(dbSession.Expiry), 0),
				TokenType:    dbSession.Tokentype,
			}
			//var tokenExtraMap = make(map[string]string) // @MarkFix doesn't work?
			//tokenExtraMap["id_token"] = auth.RawIDToken
			//restoredToken.WithExtra(tokenExtraMap)

			tokenSource := serviceCtx.Auth.TokenSource(r.Context(), restoredToken)
			newToken, err := tokenSource.Token()
			if err != nil {
				// something was wrong with the token and/or it failed to refresh
				/* @MarkFix to get here:
				   - lower "Access Token Lifespan": http://0.0.0.0:8080/admin/master/console/#/templ-realm/clients/e49746db-625e-49cd-91af-afbf8833bf95/advanced
				   - delete session from keycloak:  http://0.0.0.0:8080/admin/master/console/#/templ-realm/users/294e636d-88d7-4a62-88c2-4e9c6b0616e0/sessions
				   - wait for token to expire
				   - visit home
				*/

				queries.DeleteSession(context.Background(), dbSession.ID) // @MarkFix handle err here?
				store.DeleteSession(cookieSession, w, r)
				utils.SetContextValue(r, utils.UserCtxKey, nil)

				next.ServeHTTP(w, r)
				return
			}

			// if the token has been refreshed, verify the IDToken
			if newToken.AccessToken != restoredToken.AccessToken {
				_, err = serviceCtx.Auth.VerifyIDToken(r.Context(), newToken) // @MarkFix I need verify the nonce? According to func's code comment I do
				if err != nil {
					// @MarkFix to get here, token needs to be refreshed but IDToken not valid?

					queries.DeleteSession(context.Background(), dbSession.ID)
					store.DeleteSession(cookieSession, w, r)
					utils.SetContextValue(r, utils.UserCtxKey, nil)

					next.ServeHTTP(w, r)
					return
				}
				// @MarkFix need to update the user profile here? Shouldn't be possible for a different user to refresh the token?
				// @MarkFix maybe just assert old userID == newUserID
			}

			// if the token has been refreshed, update session
			if dbSession.Accesstoken != newToken.AccessToken {
				// @MarkFix we could update the sessionID here, to be paranoid...
				queries.UpdateSession(context.Background(), database.UpdateSessionParams{
					ID:           dbSession.ID,
					Userid:       dbSession.Userid,
					Accesstoken:  newToken.AccessToken,
					Refreshtoken: newToken.RefreshToken,
					Expiry:       newToken.Expiry.Unix(),
					Tokentype:    newToken.TokenType,
				})
			}

			// user is logged in, so set user on context for access in handlers
			loggedInUser, err := queries.GetUser(context.Background(), dbSession.Userid) // @MarkFix _should_ always return a user - add db relationship to ensure or double check here

			utils.UNUSED(err) // @MarkFix handle err?

			utils.SetContextValue(r, utils.UserCtxKey, &loggedInUser)

			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
