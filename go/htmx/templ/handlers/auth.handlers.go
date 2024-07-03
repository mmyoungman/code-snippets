package handlers

import (
	"fmt"
	"mmyoungman/templ/utils"
	"net/http"
	"text/template"

	"github.com/markbates/goth/gothic"
)

var userTemplate = `
<p><a href="/auth/logout?provider={{.Provider}}">logout</a></p>
<p>Name: {{.Name}} [{{.LastName}}, {{.FirstName}}]</p>
<p>Email: {{.Email}}</p>
<p>NickName: {{.NickName}}</p>
<p>Location: {{.Location}}</p>
<p>AvatarURL: {{.AvatarURL}} <img src="{{.AvatarURL}}"></p>
<p>Description: {{.Description}}</p>
<p>UserID: {{.UserID}}</p>
<p>AccessToken: {{.AccessToken}}</p>
<p>ExpiresAt: {{.ExpiresAt}}</p>
<p>RefreshToken: {{.RefreshToken}}</p>
`

func HandleAuthLogin(w http.ResponseWriter, r *http.Request) error {
	// try to get the user without re-authenticating
	if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, gothUser)
	} else {
		gothic.BeginAuthHandler(w, r)
	}
	return nil
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) error {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return err
	}

	// @MarkFix Instead of executing template, save user info and redirect?
	t, _ := template.New("foo").Parse(userTemplate)
	t.Execute(w, user)

	//http.Redirect(w, r, "/", http.StatusFound)
	return nil
}

func HandleAuthLogout(w http.ResponseWriter, r *http.Request) error {
	gothic.Logout(w, r)
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}
