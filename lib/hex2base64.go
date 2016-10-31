package lib

import (
	"encoding/base64"
)

// Hex2Base64Str encodes a string of hex digits into its Base64 encoded string
func Hex2Base64Str(hexBytes []byte) string {
	return base64.StdEncoding.EncodeToString(hexBytes)
}

// Base64Str2Hex decodes a Base64 string to its hex digits as a []byte
func Base64Str2Hex(base64String string) []byte  {
	bytes, err := base64.StdEncoding.DecodeString(base64String)
	CheckAbort(err)

	return bytes
}