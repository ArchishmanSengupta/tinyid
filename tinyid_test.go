package tinyid

import (
	"strings"
	"testing"
)

func TestGenerateTinyID(t *testing.T) {
	// Test with default settings
	id, err := generateTinyId(DefaultAlphabet, DefaultSize)
	if err != nil {
		t.Errorf("Error generating tiny ID: %v", err)
	}
	if len(id) != DefaultSize {
		t.Errorf("Expected ID length to be %v, got %v", DefaultSize, len(id))
	}

	// Test with custom settings
	customAlphabet := "abc123"
	customSize := 8
	id, err = generateTinyId(customAlphabet, customSize)
	if err != nil {
		t.Errorf("Error generating tiny ID with custom settings: %v", err)
	}
	if len(id) != customSize {
		t.Errorf("Expected ID length to be %v, got %v", customSize, len(id))
	}
	for _, char := range id {
		if !strings.Contains(customAlphabet, string(char)) {
			t.Errorf("ID contains invalid character: %v", char)
		}
	}

	// Test with invalid size
	_, err = generateTinyId(DefaultAlphabet, 0)
	if err == nil {
		t.Error("Expected error for invalid size, got nil")
	}
}

func TestNewTinyID(t *testing.T) {
	// Test with default settings
	id, err := NewTinyID()
	if err != nil {
		t.Errorf("Error generating new tiny ID: %v", err)
	}
	if len(id) != DefaultSize {
		t.Errorf("Expected ID length to be %v, got %v", DefaultSize, len(id))
	}
}

func BenchmarkNewTinyID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = NewTinyID()
	}
}
