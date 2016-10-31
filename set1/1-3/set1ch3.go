package main

import (
	"encoding/hex"
	"github.com/Melvillian/cryptopals/lib"
)

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexBytes, err := hex.DecodeString(hexString)
	lib.CheckAbort(err)

	_, xoredBytes, _ := lib.TryAllHexCharDecryptions(hexBytes)

	lib.PrintComparisons("Cooking MC's like a pound of bacon", string(xoredBytes))
}
