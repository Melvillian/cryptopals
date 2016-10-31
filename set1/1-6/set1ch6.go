package main

import (
	"io/ioutil"
	"github.com/Melvillian/cryptopals/lib"
	"fmt"
)

// Break repeating-key XOR
func main() {
	// read in file data and decode from base64
	rawBytes, err := ioutil.ReadFile("6.txt")
	lib.CheckAbort(err)
	hexBytes := lib.Base64Str2Hex(string(rawBytes))

	// find the most likely key size by finding which size key gives the most similar bits across consecutive blocks of bytes
	keySize := uint32(32)
	minNormHammingDist := uint32(9999999)
	var distanceSlice []uint32
	for currKeySize := uint32(2); currKeySize <= 40; currKeySize++ {
		if len(hexBytes) < 2 * 40 { panic("file data must be longer than twice the longest attempted key size!") }

		// compare 10 pairs of blocks for each size of key
		normHammingDist := uint32(0)
		for k := uint32(0); k < 10; k++ {
			hammDist, _ := lib.HammingDist(hexBytes[k * currKeySize:(k + 1) * currKeySize], hexBytes[(k + 1) * currKeySize:(k + 2) * currKeySize])
			normHammingDist += hammDist
		}
		normHammingDist = normHammingDist / currKeySize // normalize so we compare across key sizes
		distanceSlice = append(distanceSlice, normHammingDist)
		if normHammingDist < minNormHammingDist {
			minNormHammingDist = normHammingDist
			keySize = currKeySize
		}
	}

	// break the original bytes into keySize blocks
	numBlocks := uint32(len(hexBytes)) / keySize
	blockHexBytes := make([][]byte, numBlocks) // note if keySize doesn't divide len(hexBytes) we'll intentionally leave off some bytes at the end
	for i := uint32(0); i < numBlocks; i++ {
		block := make([]byte, keySize)
		for j := uint32(0); j < keySize; j++ {
			block[j] = hexBytes[(i * keySize) + j]
		}
		blockHexBytes[i] = block
	}

	// put the nth byte of each block in its own slice
	nthByteBlocks := make([][]byte, keySize)
	for _, block := range blockHexBytes {
		for in, b := range block {
			nthByteBlocks[in] = append(nthByteBlocks[in], b)
		}
	}

	var key []byte
	// for each slice of nth bytes, determine which XOR byte key is most likely to decrypt into meaningful text
	for _, nthBlock := range nthByteBlocks {
		keyByte, _, _ := lib.TryAllHexCharDecryptions(nthBlock)
		key = append(key, keyByte)
	}

	// decrypt it!
	fmt.Println(string(lib.RepeatedXor(hexBytes, key)))
}
