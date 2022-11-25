package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type Infos struct {
	Letter string
	Lives  int
}

var UsrData = Infos{
	Letter: "________",
	Lives:  10,
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {
	case "GET":
		renderTemplate(w, "hangman", UsrData)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		UsrData.Letter = r.FormValue("letter")
		renderTemplate(w, "hangman", UsrData)

	default:
		fmt.Fprintf(w, "Only GET and POST")
	}

}

func renderTemplate(w http.ResponseWriter, file string, dataFiles any) {
	t, err := template.ParseFiles("./templates/" + file + ".gohtml")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, dataFiles)
}
