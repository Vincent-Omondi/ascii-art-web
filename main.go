// main.go
package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"ascii-art/handlers"
)

func main() {
	// Use the handler function for routing
	http.HandleFunc("/", handler)

	log.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
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
