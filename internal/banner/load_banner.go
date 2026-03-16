// Package banner provides functionality for loading and processing ASCII art banners.
package banner

import (
	"os"
	"strings"
)

// LoadBanner reads a banner file from the given path, normalizes Windows-style
// line endings (CRLF) to Unix-style (LF), and returns the contents as a slice of strings.
func LoadBanner(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}
