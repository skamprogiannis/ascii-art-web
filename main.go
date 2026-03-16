// Package main is the entry point for the ascii-art-web application.
// It sets up an HTTP server to serve a web interface for generating ASCII art.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"ascii-art-web/asciiart"
)

// PageData represents the data structure passed to the HTML template for rendering.
type PageData struct {
	AsciiArt  string
	InputText string
	Banner    string
	Error     string
}

// tmpl holds the parsed HTML templates.
var tmpl *template.Template

// init parses the HTML templates on startup to avoid parsing them on every request.
func init() {
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
}

// homeHandler serves the main index page. It handles GET requests to the root URL ("/").
// It returns a 404 Not Found for any unrecognized paths and a 400 Bad Request for non-GET methods.
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

// asciiArtHandler processes form submissions to generate ASCII art.
// It handles POST requests, validates the input, and executes the template with the result.
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

// main starts the HTTP server on port 8080 and registers the application's route handlers.
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	fmt.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
