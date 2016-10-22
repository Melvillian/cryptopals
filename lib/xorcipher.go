package lib

import (
	"bytes"
	"strings"
)

var freqMap = make(map[string]float64)

// TryAllHexCharDecryptions finds the most likely character with which to XOR the given byte slice
// such that the result of the XOR will be an english sentence
func TryAllHexCharDecryptions(hexBytes []byte) byte {
	buildFrequencyMap()

	bestChar := byte(0)
	bestScore := 0.0

	for b := byte(30); b < byte(130); b++ {
		xorBytes := bytes.Repeat([]byte{b}, len(hexBytes))
		xoredBytes := XorBytes(hexBytes, xorBytes)
		score := scoreText(xoredBytes)

		if score > bestScore {
			bestScore = score
			bestChar = b
		}
	}

	return bestChar
}

// scoreText gets the score of the given text according to its word-frequency count
func scoreText(text []byte) float64 {
	var score float64 = 0.0

	for _, char := range text {
		if value, exists := freqMap[strings.ToLower(string(char))]; exists {
			score += value
		}
	}
	return score
}

// buildFrequencyMap sets word frequency percentages based on https://en.wikipedia.org/wiki/Letter_frequency
func buildFrequencyMap() {
	freqMap["a"] = 8.167
	freqMap["b"] = 1.492
	freqMap["c"] = 2.782
	freqMap["d"] = 4.253
	freqMap["e"] = 12.702
	freqMap["f"] = 2.228
	freqMap["g"] = 2.015
	freqMap["h"] = 6.094
	freqMap["i"] = 6.966
	freqMap["j"] = 0.153
	freqMap["k"] = 0.772
	freqMap["l"] = 4.025
	freqMap["m"] = 2.406
	freqMap["n"] = 6.749
	freqMap["o"] = 7.507
	freqMap["p"] = 1.929
	freqMap["q"] = 0.095
	freqMap["r"] = 5.987
	freqMap["s"] = 6.327
	freqMap["t"] = 9.056
	freqMap["u"] = 2.758
	freqMap["v"] = 0.978
	freqMap["w"] = 2.360
	freqMap["x"] = 0.150
	freqMap["y"] = 1.974
	freqMap["z"] = 0.074
	freqMap[" "] = 13.00
}
