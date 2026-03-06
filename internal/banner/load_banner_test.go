package banner

import "testing"

func TestLoadBanner_Success(t *testing.T) {
	bannerData, err := LoadBanner("../../banners/standard.txt")
	if err != nil {
		t.Fatalf("Failed to load banner: %v", err)
	}
	if len(bannerData) < 855 {
		t.Errorf("Expected at least 855 lines, got %d", len(bannerData))
	}
}

func TestLoadBanner_FileNotFound(t *testing.T) {
	_, err := LoadBanner("nonexistent.txt")
	if err == nil {
		t.Error("Expected error for missing file")
	}
}
