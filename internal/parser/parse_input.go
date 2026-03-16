// Package parser handles the extraction and validation of command-line arguments.
package parser

import (
	"fmt"
	"os"
	"strings"
)

// ParseResult represents the extracted configuration from command-line arguments,
// including the text input, the selected banner style, and optional color settings.
type ParseResult struct {
	Input       string
	Banner      string
	Color       string
	ColorSubstr string
}

// ParseInput processes os.Args to construct a ParseResult. It validates the presence
// of required arguments, parses flags (such as --color), extracts the text and banner type,
// and replaces literal "\n" sequences with actual newlines.
func ParseInput() (ParseResult, error) {
	if len(os.Args) < 2 {
		return ParseResult{}, fmt.Errorf("invalid arguments")
	}

	result := ParseResult{
		Banner: "standard",
	}

	args := os.Args[1:]
	var filteredArgs []string
	hasColor := false

	for _, arg := range args {
		if strings.HasPrefix(arg, "--color=") {
			if hasColor {
				return ParseResult{}, fmt.Errorf("invalid arguments")
			}
			hasColor = true
			colorVal := strings.TrimPrefix(arg, "--color=")
			if colorVal == "" {
				return ParseResult{}, fmt.Errorf("invalid arguments")
			}
			result.Color = colorVal
		} else {
			filteredArgs = append(filteredArgs, arg)
		}
	}

	if !hasColor && len(args) > 0 && strings.HasPrefix(args[0], "--") {
		return ParseResult{}, fmt.Errorf("invalid arguments")
	}

	if len(filteredArgs) == 0 {
		return ParseResult{}, fmt.Errorf("invalid arguments")
	}

	lastIdx := len(filteredArgs) - 1
	lastArg := filteredArgs[lastIdx]
	isBanner := false
	isBannerCandidate := false
	if strings.HasSuffix(lastArg, ".txt") {
		isBannerCandidate = true
	} else {
		lowerBanner := strings.ToLower(lastArg)
		if lowerBanner == "standard" || lowerBanner == "shadow" || lowerBanner == "thinkertoy" {
			isBannerCandidate = true
		}
	}
	if isBannerCandidate && len(filteredArgs) > 1 {
		isBanner = true
	}

	var remainingArgs []string
	if isBanner && len(filteredArgs) > 1 {
		result.Banner = lastArg
		remainingArgs = filteredArgs[:lastIdx]
	} else {
		remainingArgs = filteredArgs
	}

	if hasColor {
		if len(remainingArgs) == 1 {
			result.Input = remainingArgs[0]
			result.ColorSubstr = ""
		} else if len(remainingArgs) == 2 {
			result.ColorSubstr = remainingArgs[0]
			result.Input = remainingArgs[1]
		} else {
			return ParseResult{}, fmt.Errorf("invalid arguments")
		}
	} else {
		if len(remainingArgs) == 1 {
			result.Input = remainingArgs[0]
		} else if len(remainingArgs) == 2 && !isBanner {
			result.Input = remainingArgs[0]
			result.Banner = remainingArgs[1]
		} else if len(remainingArgs) == 2 && isBanner {
			result.Input = remainingArgs[0]
		} else {
			return ParseResult{}, fmt.Errorf("invalid arguments")
		}
	}

	result.Input = strings.ReplaceAll(result.Input, "\\n", "\n")
	return result, nil
}
