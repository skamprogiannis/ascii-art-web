package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/asciiart"
)

type PageData struct {
	AsciiArt  string
	InputText string
	Banner    string
	Error     string
}

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	data := PageData{Banner: "standard"}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	data := PageData{InputText: text, Banner: banner}

	if text != "" {
		art, err := asciiart.RenderString(text, asciiart.Options{Banner: banner})
		if err != nil {
			data.Error = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			data.AsciiArt = art
		}
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
