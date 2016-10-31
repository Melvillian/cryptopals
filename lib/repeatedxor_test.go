package lib

import (
	"testing"
	"encoding/hex"
)

func TestRepeatedXor(t *testing.T) {
	var tests = []struct {
		plaintext string;
		b []byte;
		expected string
	}{
		{"take your money", []byte{0}, hex.EncodeToString([]byte("take your money"))},
		{"Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal", []byte("ICE"), "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"},
	}
	for _, c := range tests {
		plaintextBytes := []byte(c.plaintext)

		actual := hex.EncodeToString(RepeatedXor(plaintextBytes, c.b))
		if actual != c.expected {
			t.Errorf("Expected RepeatedXor(%q, %b) to equal %q but got %q", c.plaintext, c.b, c.expected, actual)
		}
	}
}