package calc

import (
	"fmt"
	"testing"
)

func ExampleIsEven() {
	fmt.Println(IsEven(2))
	// Output: true
}

func TestIsEven(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    int
		expected bool
	}{
		{"Positive Odd", 3, false},
		{"Positive Even", 2, true},
		{"Zero", 0, true},
		{"Negative Odd", -1, false},
		{"Negative Even", -2, true},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.input)
			if got != tt.expected {
				t.Errorf("expected: %v, got: %v\n", tt.expected, got)
			}
		})
	}
}
