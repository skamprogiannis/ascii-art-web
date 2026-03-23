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

// parseANSIColor checks if the provided name matches a known standard terminal color,
// or any of the 140+ standard HTML/CSS color names.
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
		"lime":    "\033[92m",

		// Extended HTML/CSS colors mapped to Hex
		"aliceblue": "#f0f8ff", "antiquewhite": "#faebd7", "aqua": "#00ffff",
		"aquamarine": "#7fffd4", "azure": "#f0ffff", "beige": "#f5f5dc",
		"bisque": "#ffe4c4", "blanchedalmond": "#ffebcd", "blueviolet": "#8a2be2",
		"burlywood": "#deb887", "cadetblue": "#5f9ea0", "chartreuse": "#7fff00",
		"chocolate": "#d2691e", "coral": "#ff7f50", "cornflowerblue": "#6495ed",
		"cornsilk": "#fff8dc", "crimson": "#dc143c", "darkblue": "#00008b",
		"darkcyan": "#008b8b", "darkgoldenrod": "#b8860b", "darkgray": "#a9a9a9",
		"darkgreen": "#006400", "darkgrey": "#a9a9a9", "darkkhaki": "#bdb76b",
		"darkmagenta": "#8b008b", "darkolivegreen": "#556b2f", "darkorange": "#ff8c00",
		"darkorchid": "#9932cc", "darkred": "#8b0000", "darksalmon": "#e9967a",
		"darkseagreen": "#8fbc8f", "darkslateblue": "#483d8b", "darkslategray": "#2f4f4f",
		"darkslategrey": "#2f4f4f", "darkturquoise": "#00ced1", "darkviolet": "#9400d3",
		"deeppink": "#ff1493", "deepskyblue": "#00bfff", "dimgray": "#696969",
		"dimgrey": "#696969", "dodgerblue": "#1e90ff", "firebrick": "#b22222",
		"floralwhite": "#fffaf0", "forestgreen": "#228b22", "fuchsia": "#ff00ff",
		"gainsboro": "#dcdcdc", "ghostwhite": "#f8f8ff", "gold": "#ffd700",
		"goldenrod": "#daa520", "greenyellow": "#adff2f", "honeydew": "#f0fff0",
		"hotpink": "#ff69b4", "indianred": "#cd5c5c", "indigo": "#4b0082",
		"ivory": "#fffff0", "khaki": "#f0e68c", "lavender": "#e6e6fa",
		"lavenderblush": "#fff0f5", "lawngreen": "#7cfc00", "lemonchiffon": "#fffacd",
		"lightblue": "#add8e6", "lightcoral": "#f08080", "lightcyan": "#e0ffff",
		"lightgoldenrodyellow": "#fafad2", "lightgray": "#d3d3d3", "lightgreen": "#90ee90",
		"lightgrey": "#d3d3d3", "lightpink": "#ffb6c1", "lightsalmon": "#ffa07a",
		"lightseagreen": "#20b2aa", "lightskyblue": "#87cefa", "lightslategray": "#778899",
		"lightslategrey": "#778899", "lightsteelblue": "#b0c4de", "lightyellow": "#ffffe0",
		"limegreen": "#32cd32", "linen": "#faf0e6", "maroon": "#800000",
		"mediumaquamarine": "#66cdaa", "mediumblue": "#0000cd", "mediumorchid": "#ba55d3",
		"mediumpurple": "#9370db", "mediumseagreen": "#3cb371", "mediumslateblue": "#7b68ee",
		"mediumspringgreen": "#00fa9a", "mediumturquoise": "#48d1cc", "mediumvioletred": "#c71585",
		"midnightblue": "#191970", "mintcream": "#f5fffa", "mistyrose": "#ffe4e1",
		"moccasin": "#ffe4b5", "navajowhite": "#ffdead", "navy": "#000080",
		"oldlace": "#fdf5e6", "olive": "#808000", "olivedrab": "#6b8e23",
		"orangered": "#ff4500", "orchid": "#da70d6", "palegoldenrod": "#eee8aa",
		"palegreen": "#98fb98", "paleturquoise": "#afeeee", "palevioletred": "#db7093",
		"papayawhip": "#ffefd5", "peachpuff": "#ffdab9", "peru": "#cd853f",
		"plum": "#dda0dd", "powderblue": "#b0e0e6", "rebeccapurple": "#663399",
		"rosybrown": "#bc8f8f", "royalblue": "#4169e1", "saddlebrown": "#8b4513",
		"salmon": "#fa8072", "sandybrown": "#f4a460", "seagreen": "#2e8b57",
		"seashell": "#fff5ee", "sienna": "#a0522d", "silver": "#c0c0c0",
		"skyblue": "#87ceeb", "slateblue": "#6a5acd", "slategray": "#708090",
		"slategrey": "#708090", "snow": "#fffafa", "springgreen": "#00ff7f",
		"steelblue": "#4682b4", "tan": "#d2b48c", "teal": "#008080",
		"thistle": "#d8bfd8", "tomato": "#ff6347", "turquoise": "#40e0d0",
		"violet": "#ee82ee", "wheat": "#f5deb3", "whitesmoke": "#f5f5f5",
		"yellowgreen": "#9acd32",
	}

	val, ok := colors[name]
	if !ok {
		return Color{}, fmt.Errorf("unknown color")
	}

	// If the mapped value is a hex color, parse it as truecolor
	if strings.HasPrefix(val, "#") {
		return parseHexColor(val)
	}

	// Otherwise, it's an ANSI escape sequence
	return Color{Code: val}, nil
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
