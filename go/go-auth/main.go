package main

import (
	"fmt"
	"html/template"
	"log"
	"mmyoungman/go-auth/database"
	"net/http"
)

func main() {
	db, err := database.CreateDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	homeHandler := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("templates/index.html"))
		loginStatus := map[string]string{
			"LoginStatus": "Not logged in",
		}
		tmpl.Execute(w, loginStatus)
	}

	http.HandleFunc("/", homeHandler)

	fmt.Println("Webserver starting on :5501...")
	err = http.ListenAndServe(":5501", nil)
	if err != nil {
		log.Fatal(err)
	}
}
