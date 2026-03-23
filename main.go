// Package main is the entry point for the ascii-art-web application.
// It sets up an HTTP server to serve a web interface for generating ASCII art.
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"

	"ascii-art-web/asciiart"
)

// PageData represents the data structure passed to the HTML template for rendering.
type PageData struct {
	AsciiArt  template.HTML
	InputText string
	Banner    string
	Color     string
	Substr    string
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

	// Normalize text: remove Windows carriage returns (\r)
	rawText := r.FormValue("text")
	text := strings.ReplaceAll(rawText, "\r\n", "\n")
	text = strings.ReplaceAll(text, "\r", "")

	// Prevent Denial of Service (DoS) by enforcing a strict length limit
	if len(text) > 1000 {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("banner")
	colorVal := r.FormValue("color")
	substr := r.FormValue("substr")

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	data := PageData{InputText: text, Banner: banner, Color: colorVal, Substr: substr}

	// Validate that text only contains printable ASCII characters and newlines
	for _, char := range text {
		if (char < 32 || char > 126) && char != '\n' && char != '\r' {
			w.WriteHeader(http.StatusBadRequest)
			data.Error = "Invalid input: Only standard ASCII characters (32-126) are supported."
			_ = tmpl.Execute(w, data)
			return
		}
	}

	if text != "" {
		art, err := asciiart.RenderStringHTML(text, asciiart.Options{Banner: banner, Color: colorVal, Substr: substr})
		if err != nil {
			data.Error = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			data.AsciiArt = template.HTML(art)
		}
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

// main starts the HTTP server on port 8080 and registers the application's route handlers.
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Println("Server starting on http://localhost:8080")
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
