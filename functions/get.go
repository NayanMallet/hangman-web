package functions

import (
	"fmt"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")

	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website r.postfrom = %v\n", r.PostForm)
		letter := r.FormValue("letter")

		fmt.Fprintf(w, "letter = %s\n", letter)
	default:
		fmt.Fprintf(w, "Only GET and POST")
	}
}
