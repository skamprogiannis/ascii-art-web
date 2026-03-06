package parser

import (
	"os"
	"testing"
)

func TestParseInput_NoArgs(t *testing.T) {
	os.Args = []string{"cmd"}
	_, err := ParseInput()
	if err == nil {
		t.Error("Expected error for no args")
	}
}

func TestParseInput_SingleArg(t *testing.T) {
	os.Args = []string{"cmd", "Hello"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Input != "Hello" || result.Banner != "standard" {
		t.Errorf("Expected input 'Hello' and banner 'standard', got %q and %q", result.Input, result.Banner)
	}
}

func TestParseInput_TwoArgsBanner(t *testing.T) {
	os.Args = []string{"cmd", "Hello", "shadow"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Input != "Hello" || result.Banner != "shadow" {
		t.Errorf("Expected input 'Hello' and banner 'shadow', got %q and %q", result.Input, result.Banner)
	}
}

func TestParseInput_NewlineConversion(t *testing.T) {
	os.Args = []string{"cmd", "Hello\\nWorld"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Input != "Hello\nWorld" {
		t.Errorf("Expected newline conversion, got %q", result.Input)
	}
}

func TestParseInput_EmptyString(t *testing.T) {
	os.Args = []string{"cmd", ""}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Input != "" {
		t.Error("Empty string should be valid")
	}
}

func TestParseInput_ColorWholeString(t *testing.T) {
	os.Args = []string{"cmd", "--color=red", "Hello"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Color != "red" || result.ColorSubstr != "" || result.Input != "Hello" {
		t.Errorf("Expected color 'red' and input 'Hello', got color %q substr %q input %q", result.Color, result.ColorSubstr, result.Input)
	}
}

func TestParseInput_ColorSubstring(t *testing.T) {
	os.Args = []string{"cmd", "--color=red", "kit", "a king"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Color != "red" || result.ColorSubstr != "kit" || result.Input != "a king" {
		t.Errorf("Expected color 'red', substr 'kit', input 'a king', got color %q substr %q input %q", result.Color, result.ColorSubstr, result.Input)
	}
}

func TestParseInput_ColorBanner(t *testing.T) {
	os.Args = []string{"cmd", "--color=blue", "Hello", "shadow"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Color != "blue" || result.Input != "Hello" || result.Banner != "shadow" {
		t.Errorf("Expected color 'blue', input 'Hello', banner 'shadow', got color %q input %q banner %q", result.Color, result.Input, result.Banner)
	}
}

func TestParseInput_ColorSubstringBanner(t *testing.T) {
	os.Args = []string{"cmd", "--color=green", "kit", "a king", "standard"}
	result, err := ParseInput()
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result.Color != "green" || result.ColorSubstr != "kit" || result.Input != "a king" || result.Banner != "standard" {
		t.Errorf("Expected color 'green', substr 'kit', input 'a king', banner 'standard', got color %q substr %q input %q banner %q", result.Color, result.ColorSubstr, result.Input, result.Banner)
	}
}

func TestParseInput_InvalidFlagFormat(t *testing.T) {
	os.Args = []string{"cmd", "--color", "Hello"}
	_, err := ParseInput()
	if err == nil {
		t.Error("Expected error for invalid flag format")
	}
}
