package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"mmyoungman/templ/database"
	"mmyoungman/templ/handlers"
	"mmyoungman/templ/utils"
	"net/http"
	"sort"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	//"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/openidConnect"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	db := database.DBConnect()
	defer db.Close()

	router := chi.NewMux()

	// embed public dir files for prod only - so dev build hotreload works
	router.Handle("/*", public())

	// @MarkFix auth routes
	//store := sessions.NewCookieStore([]byte(utils.Getenv("KEYCLOAK_CLIENTID")))
	store := sessions.NewCookieStore([]byte("sessionkey"))
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = false

	gothic.Store = store

	openidConnect, err := openidConnect.New(
		utils.Getenv("KEYCLOAK_CLIENTID"),
		utils.Getenv("KEYCLOAK_OIDC_SECRET"),
		"http://localhost:3000/auth/openid-connect/callback",
		utils.Getenv("KEYCLOAK_DISCOVERY_URL"))
	if err != nil {
		log.Fatal("Is keycloak started? Error:\n", err)
	}
	if openidConnect != nil {
		goth.UseProviders(openidConnect)
	}

	m := map[string]string{
		"openid-connect": "OpenID Connect",
	}
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	type ProviderIndex struct {
		Providers    []string
		ProvidersMap map[string]string
	}

	providerIndex := &ProviderIndex{Providers: keys, ProvidersMap: m}

	var userTemplate = `
<p><a href="/logout/{{.Provider}}">logout</a></p>
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

	router.Get("/auth/{provider}/callback", func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		user, err := gothic.CompleteUserAuth(w, r)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		t, _ := template.New("foo").Parse(userTemplate)
		t.Execute(w, user)
	})

	router.Get("/logout/{provider}", func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		gothic.Logout(w, r)
		w.Header().Set("Location", "/")
		w.WriteHeader(http.StatusTemporaryRedirect)
	})

	router.Get("/auth/{provider}", func(w http.ResponseWriter, r *http.Request) {
		provider := chi.URLParam(r, "provider")
		r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

		// try to get the user without re-authenticating
		if gothUser, err := gothic.CompleteUserAuth(w, r); err == nil {
			t, _ := template.New("foo").Parse(userTemplate)
			t.Execute(w, gothUser)
		} else {
			gothic.BeginAuthHandler(w, r)
		}
	})

	var indexTemplate = `{{range $key,$value:=.Providers}}
    <p><a href="/auth/{{$value}}">Log in with {{index $.ProvidersMap $value}}</a></p>
{{end}}`

	router.Get("/", func(res http.ResponseWriter, req *http.Request) {
		t, _ := template.New("foo").Parse(indexTemplate)
		t.Execute(res, providerIndex)
	})

	// pages
	//router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))
	router.Get("/sign-up", handlers.Make(handlers.HandleSignUp))

	// partials
	router.Get("/test", handlers.Make(handlers.HandleTest))

	listenPort := utils.Getenv("LISTEN_PORT")
	slog.Info("Starting http server", "listenPort", listenPort)
	// @MarkFix use ListenAndServeTLS
	err = http.ListenAndServe(":"+listenPort, router)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
	fmt.Println("http server stopped")
}
