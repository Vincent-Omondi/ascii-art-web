// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"ascii-art/handlers"
)

func main() {
	// fs := http.FileServer(http.Dir("static"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handler)
	//http.HandleFunc("/ascii-art", handlers.AsciiArtHandler)

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/": handlers.IndexHandler(w, r)
	case "/ascii-art": handlers.AsciiArtHandler(w, r)
	default:
		fmt.Fprintf(w,"404 page not found")
	}
}
