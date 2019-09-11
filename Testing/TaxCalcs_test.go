// +build unit

package Testing

import (
	"testing"
)

func TestTaxCalcs(t *testing.T) {
	if calculateTax(11000) != 0 {
		t.Error("Expected no tax")
	}
}

func TestTableCalculateTax(t *testing.T) {
	var tests = []struct {
		input int
		expected float64
	} {
		{11000, 0},
		{13000, 100},
		{55000, 9500},
		{160000, 52000},
	}

	for _, test := range tests {
		if output := calculateTax(test.input); output != test.expected {
			t.Error("Test Failed: {} inputted, {} expected, received {}", test.input, test.expected, output)
		}
	}
}

//Run go test for pass/fail
//Run go test -v for tests run