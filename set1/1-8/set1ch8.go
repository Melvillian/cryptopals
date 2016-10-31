package main

import (
	"github.com/Melvillian/cryptopals/lib"
	"os"
	"bufio"
	"encoding/hex"
	"fmt"
)

func main() {
	// open file
	fi, err := os.Open("8.txt")
	defer fi.Close()
	lib.CheckAbort(err)

	// create scanner
	scanner := bufio.NewScanner(fi)

	for scanner.Scan() {
		// read line
		hexString := scanner.Text()
		hexBytes, err := hex.DecodeString(hexString)
		lib.CheckAbort(err)

		// compare each 16 byte block of the cipherText to every other 16 byte block of the cipherText
		// and check where we have matching blocks; since ECB encrypts the same plaintext to the same
		// block, if we have plaintexts which overlap then the ciphertexts will overlap as well
		if lib.CheckECB(hexBytes) {
			fmt.Println("Might have been encrypted with AEC-ECB:\n" + hexString)
		}
	}
}
