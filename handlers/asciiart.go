// handlers/asciiart.go
package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"ascii-art/asciiart"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "error 405: Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")
	fmt.Printf("%q\n", text)
	asciiChars, err := asciiart.LoadAsciiChars(banner + ".txt")
	if len(text) == 0 {
		fmt.Fprintf(w, "error 400: Bad Request")
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	str := strings.Split(text, "\r\n")
	var asciiArt string
	for _, wrd := range str {
		asciiArt += asciiart.PrintAsciiArt(wrd, asciiChars)
	}

	data := AsciiData{
		Text:   text,
		Banner: banner,
		Result: asciiArt,
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
