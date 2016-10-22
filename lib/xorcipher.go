package lib

import (
	"bytes"
	"strings"
)

var freqMap = map[string]float64{
	"a": 5, "b": 1.5, "c": 2.8, "d": 4.3, "e": 12.7, "f": 2.2, "g": 2, "h": 6.1, "i": 7.0, "j": 0.2, "k": 0.8, "l": 4, "m": 2.4,
	"n": 6.7, "o": 7.5, "p": 1.9, "q": 0.1, "r": 6, "s": 6.3, "t": 9.1, "u": 2.8, "v": 1, "w": 2.4, "x": 0.2, "y": 2, "z": 0.1,
	" ": 5,
}

// TryAllHexCharDecryptions finds the most likely character with which to XOR the given byte slice
// such that the result of the XOR will be an english sentence
func TryAllHexCharDecryptions(hexBytes []byte) (byte, []byte, float64) {
	bestChar := byte(0)
	bestScore := 0.0
	var bestXorBytes []byte

	for b := byte(30); b < byte(130); b++ {
		xorBytes := bytes.Repeat([]byte{b}, len(hexBytes))
		xoredBytes := XorBytes(hexBytes, xorBytes)
		score := ScoreText(xoredBytes)

		if score > bestScore {
			bestScore = score
			bestChar = b
			bestXorBytes = xoredBytes
		}
	}

	return bestChar, bestXorBytes, bestScore
}

// scoreText gets the score of the given text according to its word-frequency count
func ScoreText(text []byte) float64 {
	var score float64 = 0.0

	for _, char := range text {
		if value, exists := freqMap[strings.ToLower(string(char))]; exists {
			score += value
		}
	}
	return score
}
