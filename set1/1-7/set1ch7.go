package main

import (
	"crypto/aes"
	"github.com/Melvillian/cryptopals/lib"
	"io/ioutil"
	"fmt"
)

// AES in ECB mode
func main() {
	// read in file data and decode from base64
	rawBytes, err := ioutil.ReadFile("7.txt")
	lib.CheckAbort(err)
	cipherText := lib.Base64Str2Hex(string(rawBytes))

	// build cipher block using key
	key := "YELLOW SUBMARINE"
	block, err := aes.NewCipher([]byte(key))
	lib.CheckAbort(err)

	// decrypt using AESECB
	plainText := lib.AESECBModeDecrypt(cipherText, block)
	fmt.Println(string(plainText))
}
