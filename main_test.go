package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"ascii-art-web/internal/generator"
)

func newASCIIArtFormRequest(accept string, values url.Values) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if accept != "" {
		req.Header.Set("Accept", accept)
	}
	return req
}

func jsonUnmarshal(raw string, target any) error {
	return json.Unmarshal([]byte(raw), target)
}

func newExportFormRequest(values url.Values) *http.Request {
	req := httptest.NewRequest(http.MethodPost, "/export", strings.NewReader(values.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return req
}

func assertContentLength(t *testing.T, rr *httptest.ResponseRecorder) {
	t.Helper()

	want := strconv.Itoa(rr.Body.Len())
	if got := rr.Header().Get("Content-Length"); got != want {
		t.Fatalf("expected content length %q, got %q", want, got)
	}
}

func TestASCIIArtHandlerJSONSuccess(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
		"color":  {"#8b5cf6"},
		"substr": {"He"},
	}

	req := newASCIIArtFormRequest("application/json", values)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected JSON content type, got %q", ct)
	}
	assertContentLength(t, rr)

	var payload APIResponse
	if err := jsonUnmarshal(rr.Body.String(), &payload); err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}
	if payload.Error != "" {
		t.Fatalf("expected no error, got %q", payload.Error)
	}
	if payload.Banner != "standard" {
		t.Fatalf("expected banner standard, got %q", payload.Banner)
	}
	if payload.AsciiArt == "" {
		t.Fatal("expected non-empty ascii_art in JSON response")
	}
	if payload.AsciiArtHTML == "" {
		t.Fatal("expected non-empty ascii_art_html in JSON response")
	}
}

func TestASCIIArtHandlerJSONInvalidBanner(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"../../etc/passwd"},
	}

	req := newASCIIArtFormRequest("application/json", values)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}

	var payload APIResponse
	if err := jsonUnmarshal(rr.Body.String(), &payload); err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}
	if !strings.Contains(payload.Error, "Invalid banner") {
		t.Fatalf("expected invalid banner error, got %q", payload.Error)
	}
}

func TestASCIIArtHandlerJSONMissingBannerFile(t *testing.T) {
	allowedBanners["missing"] = struct{}{}
	defer delete(allowedBanners, "missing")

	values := url.Values{
		"text":   {"Hello"},
		"banner": {"missing"},
	}

	req := newASCIIArtFormRequest("application/json", values)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}

	var payload APIResponse
	if err := jsonUnmarshal(rr.Body.String(), &payload); err != nil {
		t.Fatalf("expected valid JSON response, got error: %v", err)
	}
	if payload.Error != "404 Not Found" {
		t.Fatalf("expected 404 error message, got %q", payload.Error)
	}
}

func TestASCIIArtHandlerHTMLMode(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
	}

	req := newASCIIArtFormRequest("", values)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Fatalf("expected HTML content type, got %q", ct)
	}
	assertContentLength(t, rr)

	body := rr.Body.String()
	if !strings.Contains(body, "<!DOCTYPE html>") {
		t.Fatalf("expected HTML page response, got body: %q", body)
	}
	if !strings.Contains(body, `<pre id="art-output"`) {
		t.Fatalf("expected rendered output block in response, got body: %q", body)
	}
}

func TestASCIIArtHandlerMethodValidation(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	rr := httptest.NewRecorder()
	asciiArtHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
	assertContentLength(t, rr)
}

func TestHomeHandlerNotFoundWhenTemplateUnavailable(t *testing.T) {
	oldTmpl := tmpl
	oldErr := tmplErr
	tmpl = nil
	tmplErr = errors.New("template unavailable")
	defer func() {
		tmpl = oldTmpl
		tmplErr = oldErr
	}()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	homeHandler(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %d", rr.Code)
	}
}

func TestExportHandlerTXTSuccess(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
		"color":  {"#8b5cf6"},
		"substr": {"He"},
		"format": {"txt"},
	}

	req := newExportFormRequest(values)
	rr := httptest.NewRecorder()
	exportHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "text/plain") {
		t.Fatalf("expected text content type, got %q", ct)
	}
	if cd := rr.Header().Get("Content-Disposition"); !strings.Contains(cd, `attachment; filename="ascii-art.txt"`) {
		t.Fatalf("expected txt content disposition, got %q", cd)
	}
	assertContentLength(t, rr)

	plainArt, _, err := generator.RenderBundle("Hello", generator.Options{
		Banner: "standard",
		Color:  "#8b5cf6",
		Substr: "He",
	})
	if err != nil {
		t.Fatalf("failed to render expected txt export: %v", err)
	}
	if rr.Body.String() != plainArt {
		t.Fatalf("expected txt export to match rendered art")
	}
}

func TestExportHandlerHTMLSuccess(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
		"color":  {"#8b5cf6"},
		"substr": {"He"},
		"format": {"html"},
	}

	req := newExportFormRequest(values)
	rr := httptest.NewRecorder()
	exportHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "text/html") {
		t.Fatalf("expected html content type, got %q", ct)
	}
	if cd := rr.Header().Get("Content-Disposition"); !strings.Contains(cd, `attachment; filename="ascii-art.html"`) {
		t.Fatalf("expected html content disposition, got %q", cd)
	}
	assertContentLength(t, rr)

	_, htmlArt, err := generator.RenderBundle("Hello", generator.Options{
		Banner: "standard",
		Color:  "#8b5cf6",
		Substr: "He",
	})
	if err != nil {
		t.Fatalf("failed to render expected html export: %v", err)
	}
	body := rr.Body.String()
	if !strings.Contains(body, "<!DOCTYPE html>") {
		t.Fatalf("expected html export document, got %q", body)
	}
	if !strings.Contains(body, htmlArt) {
		t.Fatalf("expected html export to contain rendered html art")
	}
}

func TestExportHandlerJSONSuccess(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
		"color":  {"#8b5cf6"},
		"substr": {"He"},
		"format": {"json"},
	}

	req := newExportFormRequest(values)
	rr := httptest.NewRecorder()
	exportHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "application/json") {
		t.Fatalf("expected json content type, got %q", ct)
	}
	if cd := rr.Header().Get("Content-Disposition"); !strings.Contains(cd, `attachment; filename="ascii-art.json"`) {
		t.Fatalf("expected json content disposition, got %q", cd)
	}
	assertContentLength(t, rr)

	var payload APIResponse
	if err := jsonUnmarshal(rr.Body.String(), &payload); err != nil {
		t.Fatalf("expected valid JSON export, got error: %v", err)
	}
	if payload.AsciiArt == "" || payload.AsciiArtHTML == "" {
		t.Fatalf("expected exported JSON to include both ascii_art and ascii_art_html")
	}
}

func TestExportHandlerInvalidFormat(t *testing.T) {
	values := url.Values{
		"text":   {"Hello"},
		"banner": {"standard"},
		"format": {"svg"},
	}

	req := newExportFormRequest(values)
	rr := httptest.NewRecorder()
	exportHandler(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Fatalf("expected status 400, got %d", rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); !strings.Contains(ct, "text/plain") {
		t.Fatalf("expected plain text error response, got %q", ct)
	}
	if !strings.Contains(rr.Body.String(), "Invalid export format") {
		t.Fatalf("expected invalid format error, got %q", rr.Body.String())
	}
	assertContentLength(t, rr)
}
