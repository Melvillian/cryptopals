package set1;

import (
	"encoding/base64"
	"encoding/hex"
)

func Hex2base64(hexString string) string {

	hexBytes, _ := hex.DecodeString(hexString)
	return base64.StdEncoding.EncodeToString(hexBytes)
}

