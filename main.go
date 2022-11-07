package main

import (
	"fmt"
	H "hangman-classic"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenue sur le jeu du pendu !")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Jouer au pendu !")
	H.Game()
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play", Play)

	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}
