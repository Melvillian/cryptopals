package lib

import "testing"

func TestXor(t *testing.T) {
	var tests = []struct {
		input1, input2, expected string
	}{
		{"1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965", "746865206b696420646f6e277420706c6179"},
		{"", "", ""},
	}
	for _, c := range tests {
		actual := Xor(c.input1, c.input2)
		if actual != c.expected {
			t.Errorf("Expected Xor(%q, %q) to equal %q but got %q", c.input1, c.input2, c.expected, actual)
		}
	}
}
