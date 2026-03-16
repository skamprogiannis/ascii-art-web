// Package color provides functionality for parsing and converting various color formats into ANSI escape codes.
package color

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Color represents a terminal color, holding its ANSI escape code sequence.
type Color struct {
	Code string
}

// ParseColor evaluates a string representation of a color and returns a Color object.
// It supports standard color names, hex codes (e.g., "#FF0000"), RGB (e.g., "rgb(255,0,0)"), and HSL formats.
func ParseColor(colorStr string) (Color, error) {
	if colorStr == "" {
		return Color{}, fmt.Errorf("empty color")
	}

	lower := strings.ToLower(colorStr)

	switch {
	case strings.HasPrefix(lower, "#"):
		return parseHexColor(colorStr)
	case strings.HasPrefix(lower, "rgb"):
		return parseRGBColor(colorStr)
	case strings.HasPrefix(lower, "hsl"):
		return parseHSLColor(colorStr)
	default:
		return parseANSIColor(lower)
	}
}

// parseANSIColor checks if the provided name matches a known standard terminal color
// and returns its corresponding ANSI escape code.
func parseANSIColor(name string) (Color, error) {
	colors := map[string]string{
		"black":   "\033[30m",
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"white":   "\033[37m",
		"orange":  "\033[38;5;208m",
		"purple":  "\033[38;5;135m",
		"pink":    "\033[38;5;205m",
		"brown":   "\033[38;5;166m",
		"gray":    "\033[90m",
		"grey":    "\033[90m",
	}

	if code, ok := colors[name]; ok {
		return Color{Code: code}, nil
	}

	return Color{}, fmt.Errorf("unknown color")
}

// parseHexColor parses a hexadecimal color string (e.g., "#RRGGBB") into a truecolor ANSI escape code.
func parseHexColor(hex string) (Color, error) {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return Color{}, fmt.Errorf("invalid hex color")
	}

	r, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return Color{}, fmt.Errorf("invalid hex color")
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return Color{}, fmt.Errorf("invalid hex color")
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return Color{}, fmt.Errorf("invalid hex color")
	}

	code := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return Color{Code: code}, nil
}

// parseRGBColor parses an RGB color string (e.g., "rgb(255, 255, 255)") into a truecolor ANSI escape code.
func parseRGBColor(rgb string) (Color, error) {
	rgb = strings.TrimPrefix(rgb, "rgb(")
	rgb = strings.TrimPrefix(rgb, "RGB(")
	rgb = strings.TrimSuffix(rgb, ")")

	parts := strings.Split(rgb, ",")
	if len(parts) != 3 {
		return Color{}, fmt.Errorf("invalid rgb color")
	}

	r, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Color{}, fmt.Errorf("invalid rgb color")
	}
	g, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	if err != nil {
		return Color{}, fmt.Errorf("invalid rgb color")
	}
	b, err := strconv.Atoi(strings.TrimSpace(parts[2]))
	if err != nil {
		return Color{}, fmt.Errorf("invalid rgb color")
	}

	if r > 255 || g > 255 || b > 255 {
		return Color{}, fmt.Errorf("invalid rgb color")
	}

	code := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return Color{Code: code}, nil
}

// parseHSLColor parses an HSL color string (e.g., "hsl(360, 100%, 100%)") into a truecolor ANSI escape code.
func parseHSLColor(hsl string) (Color, error) {
	hsl = strings.TrimPrefix(hsl, "hsl(")
	hsl = strings.TrimPrefix(hsl, "HSL(")
	hsl = strings.TrimSuffix(hsl, ")")

	parts := strings.Split(hsl, ",")
	if len(parts) != 3 {
		return Color{}, fmt.Errorf("invalid hsl color")
	}

	h, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return Color{}, fmt.Errorf("invalid hsl color")
	}
	s, err := parsePercentInt(parts[1])
	if err != nil {
		return Color{}, fmt.Errorf("invalid hsl color")
	}
	l, err := parsePercentInt(parts[2])
	if err != nil {
		return Color{}, fmt.Errorf("invalid hsl color")
	}

	if h > 360 || s > 100 || l > 100 {
		return Color{}, fmt.Errorf("invalid hsl color")
	}

	r, g, b := hslToRGB(h, s, l)
	code := fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
	return Color{Code: code}, nil
}

// hslToRGB converts Hue, Saturation, and Lightness components into Red, Green, and Blue values (0-255).
func hslToRGB(h, s, l int) (int, int, int) {
	sf := float64(s) / 100
	lf := float64(l) / 100

	c := (1 - math.Abs(2*lf-1)) * sf
	hPrime := math.Mod(float64(h), 360) / 60
	x := c * (1 - math.Abs(math.Mod(hPrime, 2)-1))
	m := lf - c/2

	var r, g, b float64

	switch {
	case hPrime >= 0 && hPrime < 1:
		r, g, b = c, x, 0
	case hPrime >= 1 && hPrime < 2:
		r, g, b = x, c, 0
	case hPrime >= 2 && hPrime < 3:
		r, g, b = 0, c, x
	case hPrime >= 3 && hPrime < 4:
		r, g, b = 0, x, c
	case hPrime >= 4 && hPrime < 5:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	return int((r + m) * 255), int((g + m) * 255), int((b + m) * 255)
}

// parsePercentInt extracts an integer from a string that may contain a trailing percentage sign.
func parsePercentInt(value string) (int, error) {
	value = strings.TrimSpace(value)
	value = strings.TrimSuffix(value, "%")
	return strconv.Atoi(value)
}
