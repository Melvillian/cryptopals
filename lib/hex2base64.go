package lib

import (
	"encoding/base64"
	"encoding/hex"
)

// Hex2base64 encodes a utf-8 string into its Base64 encoding
func Hex2base64(hexString string) string {

	hexBytes, _ := hex.DecodeString(hexString)
	return base64.StdEncoding.EncodeToString(hexBytes)
}
