package lib

import (
	"testing"
)

func TestHammingDist(t *testing.T) {
	var tests = []struct {
		input1 []byte
		input2 []byte
		expected uint32
	}{
		{[]byte("this is a test"), []byte("wokka wokka!!!"), uint32(37)},
		{[]byte(""), []byte(""), uint32(0)},
	}
	for _, c := range tests {
		actual, err := HammingDist(c.input1, c.input2)
		if err != nil { t.Errorf("Error!") }
		if actual != c.expected {
			t.Errorf("Expected HammingDist(%q, %q) to equal %d but got %v", c.input1, c.input2, c.expected, actual)
		}
	}
}