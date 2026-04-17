// Package main is the entry point for the ascii-art-web application.
// It sets up an HTTP server to serve a web interface for generating ASCII art.
package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"ascii-art-web/internal/generator"
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

// APIResponse defines the JSON response payload for /ascii-art when JSON is requested.
type APIResponse struct {
	InputText    string `json:"input_text"`
	Banner       string `json:"banner"`
	Color        string `json:"color"`
	Substr       string `json:"substr"`
	AsciiArt     string `json:"ascii_art"`
	AsciiArtHTML string `json:"ascii_art_html"`
	Error        string `json:"error"`
}

type asciiArtRequest struct {
	Text   string
	Banner string
	Color  string
	Substr string
}

const (
	maxInputLength      = 1000
	defaultExportFormat = "txt"
)

var allowedBanners = map[string]struct{}{
	"standard":   {},
	"shadow":     {},
	"thinkertoy": {},
}

var allowedExportFormats = map[string]struct{}{
	"txt":  {},
	"html": {},
	"json": {},
}

// tmpl holds the parsed HTML templates.
var (
	tmpl    *template.Template
	tmplErr error
)

// init parses the HTML templates on startup to avoid parsing them on every request.
func init() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		tmplErr = err
	}
}

// homeHandler serves the main index page. It handles GET requests to the root URL ("/").
// It returns a 404 Not Found for any unrecognized paths and a 400 Bad Request for non-GET methods.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		writePlainText(w, http.StatusNotFound, "404 Not Found")
		return
	}
	if r.Method != http.MethodGet {
		writePlainText(w, http.StatusBadRequest, "400 Bad Request")
		return
	}
	renderPage(w, http.StatusOK, PageData{Banner: "standard"})
}

// asciiArtHandler processes form submissions to generate ASCII art.
// It handles POST requests, validates the input, and executes the template with the result.
func asciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writePlainText(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	req, err := parseASCIIArtRequest(r)
	if err != nil {
		respondError(w, r, http.StatusBadRequest, asciiArtRequest{Banner: "standard"}, "400 Bad Request")
		return
	}

	if statusCode, msg := validateASCIIArtRequest(req); statusCode != http.StatusOK {
		respondError(w, r, statusCode, req, msg)
		return
	}

	pageData := toPageData(req)
	apiData := toAPIResponse(req)

	if req.Text != "" {
		plainArt, htmlArt, renderErr := generator.RenderBundle(req.Text, generator.Options{
			Banner: req.Banner,
			Color:  req.Color,
			Substr: req.Substr,
		})
		if renderErr != nil {
			statusCode, msg := classifyRenderError(renderErr)
			respondError(w, r, statusCode, req, msg)
			return
		}

		pageData.AsciiArt = template.HTML(htmlArt)
		apiData.AsciiArt = plainArt
		apiData.AsciiArtHTML = htmlArt
	}

	if wantsJSONResponse(r) {
		writeJSON(w, http.StatusOK, apiData)
		return
	}

	renderPage(w, http.StatusOK, pageData)
}

func exportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writePlainText(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	req, err := parseASCIIArtRequest(r)
	if err != nil {
		writePlainText(w, http.StatusBadRequest, "400 Bad Request")
		return
	}

	exportFormat, err := parseExportFormat(r.FormValue("format"))
	if err != nil {
		writePlainText(w, http.StatusBadRequest, err.Error())
		return
	}

	if statusCode, msg := validateASCIIArtRequest(req); statusCode != http.StatusOK {
		writePlainText(w, statusCode, msg)
		return
	}
	if req.Text == "" {
		writePlainText(w, http.StatusBadRequest, "Cannot export empty ASCII art.")
		return
	}

	plainArt, htmlArt, renderErr := generator.RenderBundle(req.Text, generator.Options{
		Banner: req.Banner,
		Color:  req.Color,
		Substr: req.Substr,
	})
	if renderErr != nil {
		statusCode, msg := classifyRenderError(renderErr)
		writePlainText(w, statusCode, msg)
		return
	}

	body, contentType, filename, buildErr := buildExportFile(exportFormat, req, plainArt, htmlArt)
	if buildErr != nil {
		writePlainText(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	writeDownload(w, http.StatusOK, contentType, filename, body)
}

// main starts the HTTP server on port 8080 and registers the application's route handlers.
func main() {
	// Serve static assets (like CSS) from the /static/ directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/ascii-art", asciiArtHandler)
	http.HandleFunc("/export", exportHandler)

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

func parseASCIIArtRequest(r *http.Request) (asciiArtRequest, error) {
	if err := r.ParseForm(); err != nil {
		return asciiArtRequest{}, err
	}

	banner := strings.ToLower(strings.TrimSpace(r.FormValue("banner")))
	if banner == "" {
		banner = "standard"
	}

	return asciiArtRequest{
		Text:   normalizeText(r.FormValue("text")),
		Banner: banner,
		Color:  strings.TrimSpace(r.FormValue("color")),
		Substr: r.FormValue("substr"),
	}, nil
}

func normalizeText(text string) string {
	normalized := strings.ReplaceAll(text, "\r\n", "\n")
	return strings.ReplaceAll(normalized, "\r", "")
}

func validateASCIIArtRequest(req asciiArtRequest) (int, string) {
	if _, ok := allowedBanners[req.Banner]; !ok {
		return http.StatusBadRequest, "Invalid banner. Allowed values: standard, shadow, thinkertoy."
	}

	if len(req.Text) > maxInputLength {
		return http.StatusBadRequest, fmt.Sprintf("Input too long: maximum %d characters.", maxInputLength)
	}

	for _, char := range req.Text {
		if char == '\n' {
			continue
		}
		if char < 32 || char > 126 {
			return http.StatusBadRequest, "Invalid input: Only standard ASCII characters (32-126) are supported."
		}
	}

	return http.StatusOK, ""
}

func classifyRenderError(err error) (int, string) {
	if errors.Is(err, fs.ErrNotExist) {
		return http.StatusNotFound, "404 Not Found"
	}
	return http.StatusInternalServerError, "500 Internal Server Error"
}

func parseExportFormat(raw string) (string, error) {
	exportFormat := strings.ToLower(strings.TrimSpace(raw))
	if exportFormat == "" {
		exportFormat = defaultExportFormat
	}
	if _, ok := allowedExportFormats[exportFormat]; !ok {
		return "", fmt.Errorf("Invalid export format. Allowed values: txt, html, json.")
	}
	return exportFormat, nil
}

func toPageData(req asciiArtRequest) PageData {
	return PageData{
		InputText: req.Text,
		Banner:    req.Banner,
		Color:     req.Color,
		Substr:    req.Substr,
	}
}

func toAPIResponse(req asciiArtRequest) APIResponse {
	return APIResponse{
		InputText: req.Text,
		Banner:    req.Banner,
		Color:     req.Color,
		Substr:    req.Substr,
	}
}

func respondError(w http.ResponseWriter, r *http.Request, statusCode int, req asciiArtRequest, message string) {
	if wantsJSONResponse(r) {
		data := toAPIResponse(req)
		data.Error = message
		writeJSON(w, statusCode, data)
		return
	}

	data := toPageData(req)
	data.Error = message
	renderPage(w, statusCode, data)
}

func wantsJSONResponse(r *http.Request) bool {
	accept := strings.ToLower(r.Header.Get("Accept"))
	return strings.Contains(accept, "application/json")
}

func writeJSON(w http.ResponseWriter, statusCode int, payload APIResponse) {
	body, err := json.Marshal(payload)
	if err != nil {
		writePlainText(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func renderPage(w http.ResponseWriter, statusCode int, data PageData) {
	if tmplErr != nil || tmpl == nil {
		writePlainText(w, http.StatusNotFound, "404 Not Found")
		return
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		writePlainText(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(body.Len()))
	w.WriteHeader(statusCode)
	_, _ = w.Write(body.Bytes())
}

func buildExportFile(exportFormat string, req asciiArtRequest, plainArt string, htmlArt string) ([]byte, string, string, error) {
	switch exportFormat {
	case "txt":
		return []byte(plainArt), "text/plain; charset=utf-8", "ascii-art.txt", nil
	case "html":
		return []byte(buildHTMLExport(req, plainArt, htmlArt)), "text/html; charset=utf-8", "ascii-art.html", nil
	case "json":
		payload := toAPIResponse(req)
		payload.AsciiArt = plainArt
		payload.AsciiArtHTML = htmlArt

		body, err := json.MarshalIndent(payload, "", "  ")
		if err != nil {
			return nil, "", "", err
		}
		return body, "application/json; charset=utf-8", "ascii-art.json", nil
	default:
		return nil, "", "", fmt.Errorf("unsupported export format")
	}
}

func buildHTMLExport(req asciiArtRequest, plainArt string, htmlArt string) string {
	artMarkup := htmlArt
	if artMarkup == "" {
		artMarkup = template.HTMLEscapeString(plainArt)
	}

	var details strings.Builder
	details.WriteString(`<li><strong>Banner:</strong> `)
	details.WriteString(template.HTMLEscapeString(req.Banner))
	details.WriteString(`</li>`)
	if req.Color != "" {
		details.WriteString(`<li><strong>Color:</strong> `)
		details.WriteString(template.HTMLEscapeString(req.Color))
		details.WriteString(`</li>`)
	}
	if req.Substr != "" {
		details.WriteString(`<li><strong>Color target:</strong> `)
		details.WriteString(template.HTMLEscapeString(req.Substr))
		details.WriteString(`</li>`)
	}

	return fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>ASCII Art Export</title>
  <style>
    body { font-family: Inter, Arial, sans-serif; margin: 2rem; background: #0f172a; color: #e2e8f0; }
    main { max-width: 1100px; margin: 0 auto; }
    h1 { margin-bottom: 0.25rem; }
    p, li { color: #cbd5e1; }
    ul { padding-left: 1.25rem; }
    pre { padding: 1.25rem; border-radius: 12px; background: #020617; border: 1px solid #1e293b; overflow: auto; white-space: pre; }
  </style>
</head>
<body>
  <main>
    <h1>ASCII Art Export</h1>
    <p><strong>Input:</strong> %s</p>
    <ul>%s</ul>
    <pre>%s</pre>
  </main>
</body>
</html>`,
		template.HTMLEscapeString(req.Text),
		details.String(),
		artMarkup,
	)
}

func writeDownload(w http.ResponseWriter, statusCode int, contentType string, filename string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf(`attachment; filename="%s"`, filename))
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}

func writePlainText(w http.ResponseWriter, statusCode int, message string) {
	body := []byte(message)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(body)))
	w.WriteHeader(statusCode)
	_, _ = w.Write(body)
}
