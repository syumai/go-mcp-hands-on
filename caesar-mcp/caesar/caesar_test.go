package caesar

import (
	"testing"
)

func TestRotN(t *testing.T) {
	tests := map[string]struct {
		input    string
		shift    int
		expected string
	}{
		"lowercase letters with positive shift": {
			input:    "abcxyz",
			shift:    3,
			expected: "defabc",
		},
		"uppercase letters with positive shift": {
			input:    "ABCXYZ",
			shift:    3,
			expected: "DEFABC",
		},
		"mixed case with positive shift": {
			input:    "AbCxYz",
			shift:    3,
			expected: "DeFaBc",
		},
		"with non-alphabetic characters": {
			input:    "Hello, World! 123",
			shift:    5,
			expected: "Mjqqt, Btwqi! 123",
		},
		"with negative shift": {
			input:    "abcxyz",
			shift:    -3,
			expected: "tuvuvw",
		},
		"with shift larger than 26": {
			input:    "abcxyz",
			shift:    29, // 29 % 26 = 3
			expected: "defabc",
		},
		"with shift equal to 26 (full rotation)": {
			input:    "abcxyz",
			shift:    26,
			expected: "abcxyz",
		},
		"empty string": {
			input:    "",
			shift:    5,
			expected: "",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			result := RotN(tt.input, tt.shift)
			if result != tt.expected {
				t.Errorf("RotN(%q, %d) = %q, expected %q",
					tt.input, tt.shift, result, tt.expected)
			}
		})
	}
}
