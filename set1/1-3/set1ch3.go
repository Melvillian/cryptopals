package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/Melvillian/cryptopals/lib"
)

func main() {
	hexString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	hexBytes, err := hex.DecodeString(hexString)
	if err != nil {
		panic(err)
	}

	probableHexByte := lib.TryAllHexCharDecryptions(hexBytes)

	xorBytes := bytes.Repeat([]byte{probableHexByte}, len(hexBytes))
	xoredBytes := lib.XorBytes(hexBytes, xorBytes)
	fmt.Println(string(xoredBytes))
}
