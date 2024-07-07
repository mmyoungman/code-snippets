package handlers

import (
	"errors"
	"fmt"
	"log"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"
	"net/url"

	"github.com/go-chi/render"
)

func HandleAuthLogin(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		state := auth.GenerateRandomState()

		session := store.GetSession(r)
		session.Values["state"] = state
		err := session.Save(r, w)
		if err != nil {
			log.Fatal("Failed to save session during login", err)
		}

		authCodeURL := authObj.AuthCodeURL(state) // @MarkFix better way to fetch authObj/Authenticator here?

		http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthCallback(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqState := r.URL.Query().Get("state")

		session := store.GetSession(r)
		state := session.Values["state"]

		if reqState != state {
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
		// @MarkFix store profile in Users table
		// @MarkFix store session info in Sessions table

		log.Println("PROFILE: ", profile)
		log.Println("UserID: ", profile["sub"])
		log.Println("Email: ", profile["email"])
		log.Println("Username: ", profile["preferred_username"])
		log.Println("Firstname: ", profile["given_name"])
		log.Println("Lastname: ", profile["family_name"])
		// given_name family_name email preferred_username
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

	state := auth.GenerateRandomState()

	session := store.GetSession(r)
	session.Values["state"] = state
	err = session.Save(r, w)
	if err != nil {
		log.Fatal("Failed to save session during logout", err)
	}

	postLogoutRedirect := fmt.Sprintf("%s%s", utils.GetPublicURL(), "/auth/logout/callback")

	parameters := url.Values{}
	parameters.Add("state", state)
	parameters.Add("id_token_hint", auth.RawIDToken)
	parameters.Add("client_id", utils.Getenv("KEYCLOAK_CLIENT_ID"))
	parameters.Add("post_logout_redirect_uri", postLogoutRedirect)	
	logoutUrl.RawQuery = parameters.Encode()

	http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)

	return nil
}

func HandleAuthLogoutCallback(w http.ResponseWriter, r *http.Request) error {
	reqState := r.URL.Query().Get("state") // @MarkFix this is probably dodgy security wise

	session := store.GetSession(r)
	state := session.Values["state"]

	if reqState != state {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//render.Status(r, http.StatusBadRequest) // @MarkFix these don't work?
		return errors.New("invalid state parameter for logout callback")
	}

	session.Values["state"] = ""
	err := session.Save(r, w)
	if err != nil {
		log.Fatal("Failed to save session during logout callback", err)
	}

	auth.AccessToken = ""
	auth.RawIDToken = ""
	auth.Profile = nil

	// @MarkFix flash up "log out success" page before redirect?

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return nil
}
