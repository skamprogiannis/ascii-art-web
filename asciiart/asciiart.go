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

	bannerName := opts.Banner
	if bannerName == "" {
		bannerName = "standard"
	}

	if bannerName != "" && !hasTxtSuffix(bannerName) {
		bannerName += ".txt"
	}

	bannerData, err := banner.LoadBanner("banners/" + bannerName)
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

	bannerName := opts.Banner
	if bannerName == "" {
		bannerName = "standard"
	}
	if !hasTxtSuffix(bannerName) {
		bannerName += ".txt"
	}

	bannerData, err := banner.LoadBanner("banners/" + bannerName)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	render.RenderArtHTML(&buf, input, bannerData, opts.Color, opts.Substr)
	return buf.String(), nil
}

// hasTxtSuffix checks whether the provided filename ends with the ".txt" extension.
func hasTxtSuffix(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[len(name)-4:] == ".txt"
}
