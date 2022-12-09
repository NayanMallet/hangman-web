package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
	"strings"
	"text/template"
)

var StartData functions.Infos

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	StartData = functions.NewGamePrep()
	StartData.WordToPrint = functions.WordToPrint(StartData.WordRune)
	StartData.Url = functions.PrintMan(StartData.Lives)
	renderTemplate(w, "hangman", StartData)
}

func Request(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/hangman" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	if StartData.Lives > 0 {
		switch r.Method {
		case "GET":
			renderTemplate(w, "hangman", StartData)
		case "POST":
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "ParseForm() err: %v", err)
				return
			}
			StartData.Propositon = strings.ToUpper(r.FormValue("letter"))
			StartData = functions.Game(StartData)
			renderTemplate(w, "hangman", StartData)

		default:
			fmt.Fprintf(w, "Only GET and POST")
		}
	} else {
		//renderTemplate(w, "game-over", StartData)
		StartData.WordToPrint = StartData.Word
		renderTemplate(w, "hangman", StartData)
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
