package banner

import (
	"os"
	"strings"
)

func LoadBanner(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	return strings.Split(content, "\n"), nil
}
