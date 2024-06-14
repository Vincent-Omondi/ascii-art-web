// main.go
package main

import (
	"log"
	"net/http"

	"ascii-art/handlers"
)

func main() {
	// Use the handler function for routing
	http.HandleFunc("/", handler)

	log.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// handler function for routing HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handlers.IndexHandler(w, r)
	case "/ascii-art":
		handlers.AsciiArtHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}
