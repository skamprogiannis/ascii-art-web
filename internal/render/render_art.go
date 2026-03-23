// Package render handles the generation and formatting of ASCII art strings, including coloring capabilities.
package render

import (
	"fmt"
	"html"
	"io"
	"strings"

	"ascii-art-web/internal/color"
)

// resetCode is the ANSI escape sequence used to reset terminal formatting.
const resetCode = "\033[0m"

// RenderArtColor writes the ASCII art representation of the input text to the given writer.
// If colorStr is provided, it applies the specified color to occurrences of colorSubstr.
func RenderArtColor(w io.Writer, input string, banner []string, colorStr string, colorSubstr string) {
	if input == "" {
		return
	}

	if strings.Trim(input, "\n") == "" {
		count := strings.Count(input, "\n")
		for i := 0; i < count; i++ {
			fmt.Fprintln(w)
		}
		return
	}

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		if line == "" {
			fmt.Fprintln(w)
		} else if colorStr == "" {
			printLine(w, line, banner)
		} else {
			printLineColor(w, line, banner, colorStr, colorSubstr)
		}
	}
}

// printLine renders a single logical line of text as 8 rows of ASCII art using the provided banner.
func printLine(w io.Writer, text string, banner []string) {
	for row := 0; row < 8; row++ {
		lineOut := ""
		for _, char := range text {
			index := int(char-32)*9 + 1 + row
			if index >= 0 && index < len(banner) {
				lineOut += banner[index]
			}
		}
		fmt.Fprintln(w, lineOut)
	}
}

// printLineColor renders a single logical line of text as 8 rows of ASCII art,
// applying color formatting to characters that match the specified substring.
func printLineColor(w io.Writer, text string, banner []string, colorStr string, colorSubstr string) {
	colorValue, err := color.ParseColor(colorStr)
	if err != nil {
		printLine(w, text, banner)
		return
	}

	coloredIndices := findColoredCharIndices(text, colorSubstr)

	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for i, char := range text {
			index := int(char-32)*9 + 1 + row
			if index >= 0 && index < len(banner) {
				charLine := banner[index]
				if coloredIndices[i] {
					sb.WriteString(colorValue.Code)
					sb.WriteString(charLine)
					sb.WriteString(resetCode)
				} else {
					sb.WriteString(charLine)
				}
			}
		}
		fmt.Fprintln(w, sb.String())
	}
}

// RenderArtHTML writes ASCII art as HTML, wrapping colored characters in <span> tags.
// Characters not in colorSubstr (or all characters when colorSubstr is empty) receive the span.
// If colorStr is empty the output is plain text (no spans).
func RenderArtHTML(w io.Writer, input string, banner []string, colorStr string, colorSubstr string) {
	if input == "" {
		return
	}

	if strings.Trim(input, "\n") == "" {
		count := strings.Count(input, "\n")
		for i := 0; i < count; i++ {
			fmt.Fprintln(w)
		}
		return
	}

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			fmt.Fprintln(w)
		} else if colorStr == "" {
			printLine(w, line, banner)
		} else {
			printLineHTML(w, line, banner, colorStr, colorSubstr)
		}
	}
}

// printLineHTML renders one logical line as 8 rows of ASCII art with HTML <span> coloring.
func printLineHTML(w io.Writer, text string, banner []string, colorStr string, colorSubstr string) {
	_, err := color.ParseColor(colorStr)
	if err != nil {
		printLine(w, text, banner)
		return
	}

	safeColor := html.EscapeString(colorStr)
	coloredIndices := findColoredCharIndices(text, colorSubstr)

	for row := 0; row < 8; row++ {
		var sb strings.Builder
		for i, char := range text {
			index := int(char-32)*9 + 1 + row
			if index >= 0 && index < len(banner) {
				charLine := html.EscapeString(banner[index])
				if coloredIndices[i] {
					sb.WriteString(`<span style="color:`)
					sb.WriteString(safeColor)
					sb.WriteString(`">`)
					sb.WriteString(charLine)
					sb.WriteString(`</span>`)
				} else {
					sb.WriteString(charLine)
				}
			}
		}
		fmt.Fprintln(w, sb.String())
	}
}

// findColoredCharIndices determines the indices of all characters in the text that are part of the target substring.
// If the substring is empty, it marks all character indices to be colored.
func findColoredCharIndices(text string, substr string) map[int]bool {
	colored := make(map[int]bool)
	if substr == "" {
		for i := 0; i < len(text); i++ {
			colored[i] = true
		}
		return colored
	}

	substrLen := len(substr)
	if substrLen == 0 {
		return colored
	}

	for i := 0; i <= len(text)-substrLen; i++ {
		if text[i:i+substrLen] == substr {
			for j := 0; j < substrLen; j++ {
				colored[i+j] = true
			}
		}
	}
	return colored
}
