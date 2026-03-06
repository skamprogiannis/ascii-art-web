package color

import "testing"

func TestParseColor_ANSIName(t *testing.T) {
	colorValue, err := ParseColor("red")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if colorValue.Code != "\033[31m" {
		t.Errorf("Expected ANSI red code, got %q", colorValue.Code)
	}
}

func TestParseColor_Hex(t *testing.T) {
	colorValue, err := ParseColor("#ff0000")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if colorValue.Code != "\033[38;2;255;0;0m" {
		t.Errorf("Expected hex red code, got %q", colorValue.Code)
	}
}

func TestParseColor_RGB(t *testing.T) {
	colorValue, err := ParseColor("rgb(0, 255, 0)")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if colorValue.Code != "\033[38;2;0;255;0m" {
		t.Errorf("Expected rgb green code, got %q", colorValue.Code)
	}
}

func TestParseColor_HSL(t *testing.T) {
	colorValue, err := ParseColor("hsl(240, 100%, 50%)")
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if colorValue.Code != "\033[38;2;0;0;255m" {
		t.Errorf("Expected hsl blue code, got %q", colorValue.Code)
	}
}

func TestParseColor_Invalid(t *testing.T) {
	_, err := ParseColor("")
	if err == nil {
		t.Error("Expected error for empty color")
	}

	_, err = ParseColor("notacolor")
	if err == nil {
		t.Error("Expected error for unknown color")
	}
}
