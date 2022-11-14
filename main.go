package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play", Play)

	functions.Openbrowser("http://localhost:8080")
	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}
