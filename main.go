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
	http.HandleFunc("/api/hangman", Request)
	http.HandleFunc("/hangman", Play)

	functions.Openbrowser("http://localhost:8080")
	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}

// idée borne d'arcade

//TODO: construct form for choose difficulty
//TODO: win/lose affichage
//TODO: add a button to restart the game
