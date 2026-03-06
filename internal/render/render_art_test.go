package render

import "testing"

func TestFindColoredCharIndices_All(t *testing.T) {
	result := findColoredCharIndices("hello", "")
	if len(result) != 5 {
		t.Errorf("Expected all characters to be colored, got %d", len(result))
	}
}

func TestFindColoredCharIndices_Substring(t *testing.T) {
	result := findColoredCharIndices("kitten kit", "kit")
	if !result[0] || !result[1] || !result[2] {
		t.Error("Expected first 'kit' to be colored")
	}
	if !result[7] || !result[8] || !result[9] {
		t.Error("Expected second 'kit' to be colored")
	}
}

func TestFindColoredCharIndices_NoMatch(t *testing.T) {
	result := findColoredCharIndices("hello", "xyz")
	if len(result) != 0 {
		t.Errorf("Expected no colored indices, got %d", len(result))
	}
}
