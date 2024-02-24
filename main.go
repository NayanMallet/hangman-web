package main

import (
	"fmt"
	"hangman-web/functions"
	"net/http"
	"os"
)

const port = ":8080"

func main() {
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/api/", RequestDifficulty)
	http.HandleFunc("/", Home)
	http.HandleFunc("/api/hangman", Request)
	http.HandleFunc("/hangman", Play)

	functions.Openbrowser("http://localhost:8080")
	fmt.Println("Serveur lancé sur le port", port)

	// Use PORT environment variable if available
	if port := os.Getenv("PORT"); port != "" {
		port = ":" + port
	}

	http.ListenAndServe(port, nil)
}

// Handler is the exported function that Vercel looks for
func Handler(w http.ResponseWriter, r *http.Request) {
	main()
}
