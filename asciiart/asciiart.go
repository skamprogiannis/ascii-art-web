// Package asciiart provides a public API for generating ASCII art text representations.
package asciiart

import (
	"ascii-art-web/internal/generator"
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
	return generator.RenderString(input, toGeneratorOptions(opts))
}

// RenderStringHTML converts input text into ASCII art and returns it as an HTML string
// with <span> tags wrapping colored characters. Safe to inject into a template as template.HTML.
func RenderStringHTML(input string, opts Options) (string, error) {
	return generator.RenderStringHTML(input, toGeneratorOptions(opts))
}

// RenderBundle converts input into both plain ASCII art text and HTML-formatted ASCII art
// using a single banner file read.
func RenderBundle(input string, opts Options) (string, string, error) {
	return generator.RenderBundle(input, toGeneratorOptions(opts))
}

func toGeneratorOptions(opts Options) generator.Options {
	return generator.Options{
		Banner: opts.Banner,
		Color:  opts.Color,
		Substr: opts.Substr,
	}
}
