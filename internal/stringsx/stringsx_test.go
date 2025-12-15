package stringsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClip(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		max      int
		expected string
	}{
		// Casos básicos
		{"string vacío", "", 5, ""},
		{"max mayor que longitud", "hello", 10, "hello"},
		{"max igual longitud", "hello", 5, "hello"},
		{"recorte normal", "hello world", 5, "hello"},

		// Casos límite/edge cases
		{"max = 0", "hello", 0, ""},
		{"max negativo", "hello", -1, ""}, // Se convierte a 0
		{"max = 1", "hello", 1, "h"},
		{"string con espacios", "hello world", 6, "hello "},
		{"unicode simple", "café", 3, "caf"}, // Cuidado con UTF-8
		{"string larga", "abcdefghijklmnopqrstuvwxyz", 10, "abcdefghij"},

		// Casos especiales
		{"max mayor pero string vacío", "", 100, ""},
		{"max 0 con string vacío", "", 0, ""},
		{"max negativo con string vacío", "", -5, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Clip(tt.input, tt.max)
			assert.Equal(t, tt.expected, result,
				"Clip(%q, %d) debería devolver %q", tt.input, tt.max, tt.expected)
		})
	}
}

// Test adicional con assert directamente
func TestClip_WithTestify(t *testing.T) {
	assert.Equal(t, "he", Clip("hello", 2))
	assert.Equal(t, "", Clip("hello", 0))
	assert.Equal(t, "", Clip("hello", -5))
	assert.Equal(t, "hello", Clip("hello", 10))
}
