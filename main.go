package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	// creates an 'http.FileServe' from the directory "assets/".
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	// The http.StripPrefix function is used to remove the "/static/" prefix from the URL path,
	// so that the file server can match the correct file in the "assets/" directory.

	http.HandleFunc("/api/", RequestDifficulty)
	http.HandleFunc("/", Home)
	http.HandleFunc("/api/hangman", Request)
	http.HandleFunc("/hangman", Play)
	// http.HandleFunc function is used to register a function to handle a specific URL path.

	functions.Openbrowser("http://localhost:8080")
	// Opens the browser at the execution of the program.
	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
	// This line starts the server and listens for incoming HTTP requests
}
