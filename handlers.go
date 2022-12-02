package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
	"text/template"
)

type Infos struct {
	Word            string
	WordRune        []rune
	WordToPrint     string
	Propositon      string
	LetterSuggested []string
	Lives           int
	Url             string
}

var StartData = Infos{
	Word:            "",
	WordRune:        nil,
	WordToPrint:     "",
	Propositon:      "",
	LetterSuggested: nil,
	Lives:           10,
	Url:             "",
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", nil)
}

func Play(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hangman" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	word, wordRune := functions.NewGamePrep()
	StartData.Lives = 10
	StartData.Word = word
	StartData.WordRune = wordRune
	StartData.WordToPrint = functions.WordToPrint(StartData.WordRune)
	StartData.Url = functions.PrintMan(StartData.Lives)
	switch r.Method {
	case "GET":
		renderTemplate(w, "hangman", StartData)

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		StartData.Propositon = r.FormValue("letter")
		renderTemplate(w, "hangman", StartData)

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
