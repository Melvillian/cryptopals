package set1

import (
	"encoding/hex"
)

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