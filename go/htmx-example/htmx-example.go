package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Film struct {
	Title    string
	Director string
}

func main() {
	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		films := map[string][]Film{
			"Films": {
				{Title: "The Godfather", Director: "Francis Ford Coppola"},
				{Title: "Blade Runner", Director: "Ridley Scott"},
				{Title: "The Thing", Director: "John Carpenter"},
			},
		}
		tmpl.Execute(w, films)
	}

	addFilmHandler := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond) // delay so Loading message is seen
		unsanitizedTitle := r.PostFormValue("title")
		unsanitizedDirector := r.PostFormValue("director")
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(
			w,
			"film-list-element",
			Film{Title: unsanitizedTitle, Director: unsanitizedDirector})
	}

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-film/", addFilmHandler)

	fmt.Print("Webserver starting on :8000...\n")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
