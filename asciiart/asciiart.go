package asciiart

import (
	"bytes"

	"ascii-art-web/internal/banner"
	"ascii-art-web/internal/render"
)

type Options struct {
	Banner string
	Color  string
	Substr string
}

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

func hasTxtSuffix(name string) bool {
	if len(name) < 4 {
		return false
	}
	return name[len(name)-4:] == ".txt"
}
