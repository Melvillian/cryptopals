package lib

import (
	"encoding/hex"
	"testing"
)

func TestScoreText(t *testing.T) {
	var tests = []struct {
		input1   string
		expected float64
	}{
		{"", 0.0},
		{"50", 1.9},
		{"00", 0.0},
		{"504f", 9.4},
	}
	for _, c := range tests {
		bytes, err := hex.DecodeString(c.input1)
		if err != nil {
			panic(err)
		}
		actual := ScoreText(bytes)
		if actual != c.expected {
			t.Errorf("Expected ScoreText(%q) to equal %f but got %f", c.input1, c.expected, actual)
		}
	}
}
