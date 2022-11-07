package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/play", Play)

	fmt.Println("Serveur lancé sur le port", port)
	http.ListenAndServe(port, nil)
}
