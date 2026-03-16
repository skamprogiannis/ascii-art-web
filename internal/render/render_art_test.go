// Package render handles the generation and formatting of ASCII art from text.
package render

import "testing"

// TestFindColoredCharIndices_All verifies that when no specific substring is provided,
// all characters in the input string are marked for coloring.
func TestFindColoredCharIndices_All(t *testing.T) {
	result := findColoredCharIndices("hello", "")
	if len(result) != 5 {
		t.Errorf("Expected all characters to be colored, got %d", len(result))
	}
}

// TestFindColoredCharIndices_Substring checks that only the characters belonging to
// the exact matching substring are marked for coloring.
func TestFindColoredCharIndices_Substring(t *testing.T) {
	result := findColoredCharIndices("kitten kit", "kit")
	if !result[0] || !result[1] || !result[2] {
		t.Error("Expected first 'kit' to be colored")
	}
	if !result[7] || !result[8] || !result[9] {
		t.Error("Expected second 'kit' to be colored")
	}
}

// TestFindColoredCharIndices_NoMatch ensures that no characters are marked for coloring
// when the target substring does not exist within the input text.
func TestFindColoredCharIndices_NoMatch(t *testing.T) {
	result := findColoredCharIndices("hello", "xyz")
	if len(result) != 0 {
		t.Errorf("Expected no colored indices, got %d", len(result))
	}
}
