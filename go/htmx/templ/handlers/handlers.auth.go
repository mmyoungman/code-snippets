package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/store"
	"mmyoungman/templ/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/render"
	"github.com/google/uuid"
)

func HandleAuthLogin(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		state := auth.GenerateRandomState()

		session := store.GetSession(r)
		session.Values["state"] = state

		// get referrer URL/path so can redirect user to page they were previously on after login
		referrer := r.Header.Get("Referer") // @MarkFix could get this ourselves to prevent future browser change issues?
		if (strings.HasPrefix(referrer, "/") || strings.HasPrefix(referrer, utils.GetPublicURL())) {
			session.Values["referrer_path"] = referrer
		}

		err := session.Save(r, w)
		if err != nil {
			log.Fatal("Failed to save session during login", err)
		}

		authCodeURL := authObj.AuthCodeURL(state) // @MarkFix better way to fetch authObj/Authenticator here?

		http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthCallback(authObj *auth.Authenticator, db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqState := r.URL.Query().Get("state")

		session := store.GetSession(r)

		state := session.Values["state"]
		referrerPath := session.Values["referrer_path"]

		session.Values["state"] = nil
		session.Values["referrer_path"] = nil

		err := session.Save(r, w)
		if err != nil {
			log.Fatal("Failed to save session during login callback - ", err)
		}

		if reqState != state {
			render.Status(r, http.StatusBadRequest)
			return errors.New("invalid state parameter")
		}

		reqCode := r.URL.Query().Get("code")

		token, err := authObj.Exchange(r.Context(), reqCode)
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
		if err = idToken.Claims(&profile); err != nil {
			render.Status(r, http.StatusInternalServerError)
			return err
		}

		newSessionID := uuid.New().String()
		userID := profile["sub"].(string)
		cookieSession := store.GetSession(r)
		cookieSession.Values["session_id"] = newSessionID
		cookieSession.Values["user_id"] = userID
		err = cookieSession.Save(r, w)
		if err != nil {
			log.Fatal("Failed to save cookie session")
		}

		database.InsertSession(db, newSessionID, userID, token.AccessToken, token.RefreshToken, token.Expiry.Unix(), token.TokenType)

		auth.RawIDToken = token.Extra("id_token").(string)
		auth.Profile = profile
		// @MarkFix store profile in Users table

		//fmt.Printf("Profile string: '%s'\n", session.Values["profile"])
		//log.Println("PROFILE: ", profile)
		//log.Println("UserID: ", profile["sub"])
		//log.Println("Email: ", profile["email"])
		//log.Println("Username: ", profile["preferred_username"])
		//log.Println("Firstname: ", profile["given_name"])
		//log.Println("Lastname: ", profile["family_name"])
		//log.Println("RawIDToken: ", auth.RawIDToken)

		if referrerPath == nil || referrerPath.(string) == "" { // @MarkFix review the referrer thing entirely
			referrerPath = "/"
		}
		http.Redirect(w, r, referrerPath.(string), http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthLogout(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// @MarkFix does this log the user out of the idp entirely, or just for this site? i.e. would this work with google/facebook?
		logoutUrl, err := url.Parse(authObj.EndSessionURL)
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
}

func HandleAuthLogoutCallback(w http.ResponseWriter, r *http.Request) error {
	reqState := r.URL.Query().Get("state") // @MarkFix this ok in terms of security?

	session := store.GetSession(r)
	state := session.Values["state"]
	session.Values["state"] = nil
	err := session.Save(r, w)
	if err != nil {
		log.Fatal("Failed to save session during logout callback", err)
	}


	if reqState != state {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		//render.Status(r, http.StatusBadRequest) // @MarkFix these don't work?
		return errors.New("invalid state parameter for logout callback")
	}

	session.Values["session_id"] = nil
	session.Values["user_id"] = nil

	err = session.Save(r, w)
	if err != nil {
		log.Fatal("Failed to save session during logout callback", err)
	}

	auth.RawIDToken = ""
	auth.Profile = nil

	// @MarkFix flash up "log out success" page before redirect?

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	return nil
}
