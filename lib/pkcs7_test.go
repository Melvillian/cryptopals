package lib

import (
	"testing"
	"reflect"
)

func TestPadWithPKCS7(t *testing.T) {
	var tests = []struct {
		data []byte
		blockSize int
		expected []byte
	}{
		{[]byte("YELLOW SUBMARINE"), 20, []byte("YELLOW SUBMARINE\x04\x04\x04\x04")},
		{[]byte(""), 10, []byte("\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a\x0a")},
		{[]byte("abc"), 3, []byte("abc\x03\x03\x03")},
	}
	for _, c := range tests {
		actual := PadWithPKCS7(c.data, c.blockSize)
		if !reflect.DeepEqual(c.expected, actual) {
			t.Errorf("Expected PadWithPKCS7(%q, %q) to equal %q but got %q", c.data, c.blockSize, c.expected, actual)
		}
	}
}