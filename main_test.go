package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
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
