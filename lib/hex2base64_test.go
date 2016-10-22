package lib

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"1c0111001f010100061a024b53535009181c", "HAERAB8BAQAGGgJLU1NQCRgc="},
		{"", ""},
	}
	for _, c := range tests {
		fmt.Println(Hex2base64(c.input))
		actual := Hex2base64(c.input)
		if actual != c.expected {
			t.Errorf("Expected Hex2base64(%q) to equal %q but got %q", c.input, c.expected, actual)
		}
	}
}
