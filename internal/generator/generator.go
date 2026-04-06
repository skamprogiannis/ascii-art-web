// Package generator provides shared ASCII art generation helpers for application code.
package generator

import (
	"bytes"

	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/render"
)

// Options contains configuration settings for generating ASCII art.
type Options struct {
	Banner string
	Color  string
	Substr string
}

// RenderString converts the provided input string into plain ASCII art text.
func RenderString(input string, opts Options) (string, error) {
	if input == "" {
		return "", nil
	}

	bannerData, err := loadBannerData(opts.Banner)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	render.RenderArtColor(&buf, input, bannerData, opts.Color, opts.Substr)
	return buf.String(), nil
}

// RenderStringHTML converts the provided input string into HTML-formatted ASCII art.
func RenderStringHTML(input string, opts Options) (string, error) {
	if input == "" {
		return "", nil
	}

	bannerData, err := loadBannerData(opts.Banner)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	render.RenderArtHTML(&buf, input, bannerData, opts.Color, opts.Substr)
	return buf.String(), nil
}

// RenderBundle converts input into both plain ASCII art text and HTML-formatted ASCII art.
func RenderBundle(input string, opts Options) (string, string, error) {
	if input == "" {
		return "", "", nil
	}

	bannerData, err := loadBannerData(opts.Banner)
	if err != nil {
		return "", "", err
	}

	var plain bytes.Buffer
	render.RenderArtColor(&plain, input, bannerData, "", "")

	var html bytes.Buffer
	render.RenderArtHTML(&html, input, bannerData, opts.Color, opts.Substr)

	return plain.String(), html.String(), nil
}

func loadBannerData(bannerName string) ([]string, error) {
	resolved := bannerName
	if resolved == "" {
		resolved = "standard"
	}
	if !hasTxtSuffix(resolved) {
		resolved += ".txt"
	}
	return banner.LoadBanner("banners/" + resolved)
}

func hasTxtSuffix(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[len(name)-4:] == ".txt"
}
