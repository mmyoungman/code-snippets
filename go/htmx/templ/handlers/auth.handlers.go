package handlers

import (
	"errors"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/utils"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)

func HandleAuthLogin(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		state, err := auth.GenerateRandomState()
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return errors.New("could not generate random state for auth login")
		}

		auth.State = state

		authCodeURL := authObj.AuthCodeURL(auth.State) // @MarkFix better way to fetch authObj/Authenticator here?

		http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthCallback(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqState := r.URL.Query().Get("state")
		if reqState != auth.State {
			render.Status(r, http.StatusBadRequest)
			return errors.New("invalid state parameter")
		}

		token, err := authObj.Exchange(r.Context(), r.URL.Query().Get("code"))
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			//return errors.New("Failed to convert authorization code into a token")
			return err
		}

		idToken, err := authObj.VerifyIDToken(r.Context(), token)
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

		auth.RawIDToken = token.Extra("id_token").(string)

		auth.AccessToken = token.AccessToken
		auth.RefreshToken = token.RefreshToken
		auth.TokenType = token.TokenType
		auth.Expiry = token.Expiry

		auth.Profile = profile

		//log.Println("PROFILE: ", profile)
		//log.Println("RawIDToken: ", auth.RawIDToken)

		// @MarkFix redirect back to page they logged in from or to a logged in user page
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthLogout(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix does this log the user out of the idp entirely, or just for this site? i.e. would this work with google/facebook?
	logoutUrl, err := url.Parse(auth.EndSessionURL)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		return err
	}

	state, err := auth.GenerateRandomState()
	if err != nil {
		return err
	}
	auth.State = state

	parameters := url.Values{}
	parameters.Add("state", auth.State)
	parameters.Add("id_token_hint", auth.RawIDToken)
	parameters.Add("client_id", utils.Getenv("KEYCLOAK_CLIENT_ID"))
	parameters.Add("post_logout_redirect_uri", utils.Getenv("PUBLIC_HOST") + ":" + utils.Getenv("PUBLIC_PORT") + "/auth/logout/callback")	
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)

	return nil
}

func HandleAuthLogoutCallback(w http.ResponseWriter, r *http.Request) error {
	reqState := r.URL.Query().Get("state") // @MarkFix this is probably dodgy security wise
	if reqState != auth.State {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//render.Status(r, http.StatusBadRequest) // @MarkFix these don't work?
		return errors.New("invalid state parameter for logout callback")
	}

	auth.State = ""
	auth.AccessToken = ""
	auth.RawIDToken = ""
	auth.Profile = nil

	// @MarkFix flash up "log out success" page before redirect?

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return nil
}
