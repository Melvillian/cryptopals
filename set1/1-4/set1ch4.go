package main

import (
	"bufio"
	"encoding/hex"
	"github.com/Melvillian/cryptopals/lib"
	"os"
	"strings"
)

// Detect single-character XOR
func main() {
	// open file
	fi, err := os.Open("4.txt")
	defer fi.Close()
	lib.CheckAbort(err)

	// create scanner
	scanner := bufio.NewScanner(fi)

	bestScore := float64(0)
	decryptedBestLine := ""
	for scanner.Scan() {
		// read line
		hexString := scanner.Text()
		hexBytes, err := hex.DecodeString(hexString)
		lib.CheckAbort(err)

		// get best guess at single XOR byte, and get that score
		_, bestXorBytes, score := lib.TryAllHexCharDecryptions(hexBytes)
		if score > bestScore {
			bestScore = score
			decryptedBestLine = string(bestXorBytes)
		}

	}
	lib.PrintComparisons("Now that the party is jumping", strings.Trim(decryptedBestLine, "\n"))
}
