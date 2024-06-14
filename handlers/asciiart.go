// handlers/asciiart.go
package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"ascii-art/asciiart"
)

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// check if the method is not POST
	if r.Method != http.MethodPost {
		http.Error(w, "error 405: Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// receive tesxt and banner from text field and banner option
	text := r.FormValue("text")
	banner := r.FormValue("banner")
	asciiChars, err := asciiart.LoadAsciiChars(banner + ".txt")

	// check if the text field is empty
	if len(text) == 0 {
		http.Error(w, "error 400: Bad Request", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// split the text using  newline and carriage return and creates the art and appends it to asciiArt
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

	// Execute the data and send it to tmpl
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
