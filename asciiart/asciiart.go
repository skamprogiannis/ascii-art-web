// Package asciiart provides a public API for generating ASCII art text representations.
package asciiart

import (
	"bytes"

	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/render"
)

// Options contains configuration settings for generating ASCII art, such as the banner style and color.
type Options struct {
	Banner string
	Color  string
	Substr string
}

// RenderString converts the provided input string into its ASCII art equivalent
// based on the given Options, and returns the resulting string.
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

// RenderStringHTML converts input text into ASCII art and returns it as an HTML string
// with <span> tags wrapping colored characters. Safe to inject into a template as template.HTML.
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

// RenderBundle converts input into both plain ASCII art text and HTML-formatted ASCII art
// using a single banner file read.
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

// hasTxtSuffix checks whether the provided filename ends with the ".txt" extension.
func hasTxtSuffix(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[len(name)-4:] == ".txt"
}
