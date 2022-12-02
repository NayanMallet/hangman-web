package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/hangman", Play)

	functions.Openbrowser("http://localhost:8080")
	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}

// idée borne d'arcade

// send get/post request
