package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"mmyoungman/templ/auth"
	"mmyoungman/templ/database"
	"mmyoungman/templ/database/jet/model"
	"mmyoungman/templ/store"
	"mmyoungman/templ/structs"
	"mmyoungman/templ/utils"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-chi/render"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func HandleAuthLogin(authObj *auth.Authenticator) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {

		session := store.GetSession(r)

		state := auth.GenerateRandomState()
		session.Values["state_login"] = state

		// get referrer URL/path so can redirect user to page they were previously on after login
		referrer := r.Header.Get("Referer")
		if strings.HasPrefix(referrer, "/") || strings.HasPrefix(referrer, utils.GetPublicURL()) {
			session.Values["referrer_path"] = referrer
		}

		pkceVerifier := oauth2.GenerateVerifier()
		session.Values["pkce_verifier"] = pkceVerifier

		store.SaveSession(session, w, r)

		pkceChallengeOption := oauth2.S256ChallengeOption(pkceVerifier)

		authCodeURL := authObj.AuthCodeURL(state, pkceChallengeOption)

		http.Redirect(w, r, authCodeURL, http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthCallback(serviceCtx *structs.ServiceCtx) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqState := r.URL.Query().Get("state")

		session := store.GetSession(r)

		state := session.Values["state_login"]
		pkceVerifier := session.Values["pkce_verifier"]
		referrerPath := session.Values["referrer_path"]

		session.Values["state_login"] = nil
		session.Values["pkce_verifier"] = nil
		session.Values["referrer_path"] = nil

		store.SaveSession(session, w, r)

		if reqState != state {
			render.Status(r, http.StatusBadRequest)
			return errors.New("invalid state parameter")
		}

		reqCode := r.URL.Query().Get("code")
		pkceVerifierOption := oauth2.VerifierOption(pkceVerifier.(string))

		token, err := serviceCtx.Auth.Exchange(r.Context(), reqCode, pkceVerifierOption)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			//return errors.New("Failed to convert authorization code into a token")
			return err
		}

		idToken, err := serviceCtx.Auth.VerifyIDToken(r.Context(), token)
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

		userID := profile["sub"].(string)

		// insert/update db user first
		rawIDToken := token.Extra("id_token").(string)
		user := database.GetUser(serviceCtx.Db, profile["sub"].(string))
		if user == nil {
			database.InsertUser(serviceCtx.Db, &model.User{ // @MarkFix make stateless?
				ID:         userID,
				Username:   profile["preferred_username"].(string),
				Email:      profile["email"].(string),
				FirstName:  profile["given_name"].(string),
				LastName:   profile["family_name"].(string),
				RawIDToken: rawIDToken,
			})
		} else {
			database.UpdateUser(serviceCtx.Db, &model.User{
				ID:         userID,
				Username:   profile["preferred_username"].(string),
				Email:      profile["email"].(string),
				FirstName:  profile["given_name"].(string),
				LastName:   profile["family_name"].(string),
				RawIDToken: rawIDToken,
			})
		}

		// then insert new db session
		newSessionID := uuid.NewString()
		database.InsertSession(serviceCtx.Db, newSessionID, userID, token.AccessToken, token.RefreshToken, token.Expiry.Unix(), token.TokenType)

		// then update cookie session
		cookieSession := store.GetSession(r)
		cookieSession.Values["session_id"] = newSessionID
		store.SaveSession(cookieSession, w, r)

		if referrerPath == nil || referrerPath.(string) == "" {
			referrerPath = "/"
		}
		http.Redirect(w, r, referrerPath.(string), http.StatusTemporaryRedirect)
		return nil
	}
}

func HandleAuthLogout(serviceCtx *structs.ServiceCtx) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		// @MarkFix does this log the user out of the idp entirely, or just for this site? i.e. would this work with google/facebook?

		var user *model.User // @MarkFix clean this up
		userUntyped := r.Context().Value(utils.ReqUserCtxKey)
		if userUntyped == nil {
			// @MarkFix anything else to do here? Clean up session or something?
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		} else {
			user = userUntyped.(*model.User)
		}

		cookieSession := store.GetSession(r) // @MarkFix don't need this now? Get user from context

		// get referrer URL/path so can redirect user to page they were previously on after logout
		referrer := r.Header.Get("Referer") // @MarkFix could get this ourselves to prevent future browser change issues?
		if strings.HasPrefix(referrer, "/") || strings.HasPrefix(referrer, utils.GetPublicURL()) {
			cookieSession.Values["referrer_path"] = referrer
		}

		state := auth.GenerateRandomState()
		cookieSession.Values["state_logout"] = state

		store.SaveSession(cookieSession, w, r)

		// construct redirect URL + query params
		logoutUrl, err := url.Parse(serviceCtx.Auth.EndSessionURL)
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			return err
		}

		postLogoutRedirect := fmt.Sprintf("%s%s", utils.GetPublicURL(), "/auth/logout/callback")

		parameters := url.Values{}
		parameters.Add("state", state)
		parameters.Add("id_token_hint", user.RawIDToken)
		parameters.Add("client_id", utils.Getenv("KEYCLOAK_CLIENT_ID"))
		parameters.Add("post_logout_redirect_uri", postLogoutRedirect)
		logoutUrl.RawQuery = parameters.Encode()

		http.Redirect(w, r, logoutUrl.String(), http.StatusTemporaryRedirect)

		return nil
	}
}

func HandleAuthLogoutCallback(db *sql.DB) HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		reqState := r.URL.Query().Get("state") // @MarkFix this ok in terms of security?

		session := store.GetSession(r)

		state := session.Values["state_logout"]
		referrerPath := session.Values["referrer_path"]

		session.Values["state_logout"] = nil
		session.Values["referrer_path"] = nil

		store.SaveSession(session, w, r)

		if reqState != state {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			//render.Status(r, http.StatusBadRequest) // @MarkFix these don't work?
			return errors.New("invalid state parameter for logout callback")
		}

		sessionID := session.Values["session_id"].(string)
		database.DeleteSession(db, sessionID)
		store.DeleteSession(session, w, r)
		// @MarkFix delete user here?

		// @MarkFix flash up "log out success" page before redirect?

		http.Redirect(w, r, referrerPath.(string), http.StatusTemporaryRedirect)
		return nil
	}
}
