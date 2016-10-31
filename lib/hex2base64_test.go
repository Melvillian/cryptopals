package lib

import (
	"testing"
	"encoding/hex"
)

func TestBase64Str2Hex(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"},
		{"HAERAB8BAQAGGgJLU1NQCRgc", "1c0111001f010100061a024b53535009181c"},
		{"", ""},
	}
	for _, c := range tests {
		actual := hex.EncodeToString(Base64Str2Hex(c.input))
		if actual != c.expected {
			t.Errorf("Expected Base64Str2Hex(%q) to equal %q but got %q", c.input, c.expected, actual)
		}
	}
}

func TestHex2Base64(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"1c0111001f010100061a024b53535009181c", "HAERAB8BAQAGGgJLU1NQCRgc"},
		{"", ""},
	}
	for _, c := range tests {
		hexBytes, _ := hex.DecodeString(c.input)
		actual := Hex2Base64Str(hexBytes)
		if actual != c.expected {
			t.Errorf("Expected Hex2Base64Str(%q) to equal %q but got %q", c.input, c.expected, actual)
		}
	}
}