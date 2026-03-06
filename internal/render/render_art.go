package render

import (
	"fmt"
	"io"
	"strings"

	"ascii-art-web/internal/color"
)

const resetCode = "\033[0m"

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
