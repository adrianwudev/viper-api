package routes

import (
	"fmt"
	"io"
	"net/http"
)

func Start() {
	http.HandleFunc("/", greet)

	http.ListenAndServe(":8080", nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Println("There's a gopher commin' in")
	io.WriteString(w, "Hello gopher.\n")
}
