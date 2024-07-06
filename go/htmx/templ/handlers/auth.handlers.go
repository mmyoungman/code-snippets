package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/utils"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)


func HandleAuthLogin(w http.ResponseWriter, r *http.Request) error {
	state, err := generateRandomState()
	if err != nil {
		// @MarkFix log err
		render.Status(r, http.StatusInternalServerError)
	}

	auth.State = state

	authCodeURL := auth.Auth.AuthCodeURL(state) // @MarkFix better way to pass Authenticator here?

	http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
	return nil
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) error {
	reqState := r.URL.Query().Get("state")
	if reqState != auth.State {
		render.Status(r, http.StatusBadRequest)
		return errors.New("invalid state parameter")
	}

	token, err := auth.Auth.Exchange(r.Context(), r.URL.Query().Get("code"))
	if err != nil {
		render.Status(r, http.StatusUnauthorized)
		//return errors.New("Failed to convert authorization code into a token")
		return err
	}

	idToken, err := auth.Auth.VerifyIDToken(r.Context(), token)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		//return errors.New("Failed to verify ID Token")
		return err
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		render.Status(r, http.StatusInternalServerError)
		return err
	}

	log.Print(profile)

	auth.AccessToken = token.AccessToken
	auth.Profile = profile

	// @MarkFix redirect back to page they logged in from or to a logged in user page
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return nil
}

func HandleAuthLogout(w http.ResponseWriter, r *http.Request) error {
	logoutUrl, err := url.Parse(utils.Getenv("KEYCLOAK_URL") + "/v2/logout") // @MarkFix almost definitely wrong
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return err
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + r.Host)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return err
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", utils.Getenv("KEYCLOAK_CLIENT_ID"))
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)
	return nil
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

