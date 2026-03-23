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
	AsciiArt  template.HTML // The generated ASCII art (safe HTML)
	InputText string        // The original text submitted by the user
	Banner    string        // The selected banner style (standard, shadow, thinkertoy)
	Color     string        // The requested color (hex, rgb, or name)
	Substr    string        // The specific substring to color
	Error     string        // Any error message to display to the user
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

	// Parse the incoming form data from the POST request body
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
		w.WriteHeader(http.StatusBadRequest)
		_ = tmpl.Execute(w, PageData{
			Banner: "standard",
			Error:  "Nice try, hacker! 🕵️‍♂️ Did you really think you could crash the server by modifying the HTML limits? My backend is bulletproof! 🛡️ (Max 1000 characters)",
		})
		return
	}

	// Extract remaining form values
	banner := r.FormValue("banner")
	colorVal := r.FormValue("color")
	substr := r.FormValue("substr")

	// Strict banner validation to prevent Path Traversal attacks (e.g., banner=../../../etc/passwd)
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		w.WriteHeader(http.StatusBadRequest)
		_ = tmpl.Execute(w, PageData{
			Banner: "standard",
			Error:  "Hold on! 🛑 That banner doesn't exist. Trying to inject some path traversal? Not today! 🚫",
		})
		return
	}

	// Initialize the data payload with the user's input state
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

	// If the text is valid and not empty, generate the stylized ASCII art
	if text != "" {
		art, err := asciiart.RenderStringHTML(text, asciiart.Options{Banner: banner, Color: colorVal, Substr: substr})
		if err != nil {
			data.Error = err.Error()
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			data.AsciiArt = template.HTML(art)
		}
	}

	// Render the HTML template, injecting the generated art or any error messages
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	}
}

// main starts the HTTP server on port 8080 and registers the application's route handlers.
func main() {
	// Serve static assets (like CSS) from the /static/ directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)

	// Configure a robust HTTP server with strict timeouts to prevent Slowloris connection attacks
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
