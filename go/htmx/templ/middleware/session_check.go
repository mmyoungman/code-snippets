package middleware

import (
	"fmt"
	"log"
	"mmyoungman/templ/auth"
	"net/http"

	"golang.org/x/oauth2"
)

// Checks whether the user has a session. If they do, validate it and/or refresh token
func SessionCheck(authObj *auth.Authenticator) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			var count int = 1 // for testing
			// @MarkFix If someone changed Token.Expiry to far in the future, we'd curerntly 
			// never revalidate the user and they'd stay logged in permanently? So maybe
			// VerifyIDToken every time?
			if auth.Token != nil && !auth.Token.Valid() {
				var err error
				restoredToken := &oauth2.Token{
					AccessToken:  auth.Token.AccessToken,
					RefreshToken: auth.Token.RefreshToken,
					Expiry:       auth.Token.Expiry,
					TokenType:    auth.Token.TokenType,
				}
				tokenSource := authObj.TokenSource(r.Context(), restoredToken)
				auth.Token, err = tokenSource.Token()
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
				_, err = authObj.VerifyIDToken(r.Context(), auth.Token) // @MarkFix do we bother with this? - and I need verify the nonce?
				if err != nil {
					log.Fatal("Error verifying id token: ", err)
					// @MarkFix remove any session? maybe not - could be used to invalidate a real session?
					// @MarkFix redirect to Home?
				}
				fmt.Println("\nSessionCheck middleware - access token refreshed. Count: ", count, auth.Token.Expiry)
				count++

				// @MarkFix if the token has changed, save stuff
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
