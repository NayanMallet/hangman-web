package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home")
}

func Play(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "templates/hangman.page.tmpl")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostForm)
		letter := r.FormValue("letter")

		fmt.Fprintf(w, "letter = %s\n", letter)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, nil)
}
