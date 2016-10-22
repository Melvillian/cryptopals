package lib

import (
	"encoding/hex"
)

// Xor expects two hexadecimal strings, and computes the logical XOR of the their byte slices
func Xor(str1, str2 string) string {
	str1Hex, _ := hex.DecodeString(str1)
	str2Hex, _ := hex.DecodeString(str2)


	n := len(str1Hex) // assume they're both the same length
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = str1Hex[i] ^ str2Hex[i]
	}
	return hex.EncodeToString(b)
}

// Xor expects two byte slices, and computes their logical XOR
func XorBytes(bytes1, bytes2 []byte) []byte {
	n := len(bytes1) // assume they're both the same length
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		b[i] = bytes1[i] ^ bytes2[i]
	}
	return b
}