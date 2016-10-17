package set1

import "fmt"

var freqMap map[rune]float32 = make(map[rune]float32)


func TryAllHexCharDecryptions(hexString string) byte {
	allHex := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	scores := make(map[byte]float32)

	for _, b := range allHex {
		extended := b * len(hexString)
		xoredText := Xor(hexString, extended)
		score := scoreText(xoredText)
		scores[b] = score
	}

	var currMax float32 = 0.0
	var currMaxByte byte = 0
	for index, score := range scores {
		if (score > currMax) {
			currMax = score
			currMaxByte = index

		}
	}
	return currMaxByte
}

// make a string array all with the same values and a given length
func stringExtender(char string, length int) string {
	charBytes := []byte(char)
	var extended = make([]byte, length)
	for x := 0; x < length; x++ {
		extended = append(extended, charBytes[0])
	}
	return string(extended)
}

// get the score of the given text according to its word-frequency count
func scoreText(text string) float32 {
	buildFrequencyMap()
	var total float32 =  0.0
	for _, word := range text {
		if freqMap[word] != 0.0 {
			total += freqMap[word]
		}
	}
	return total
}

// set word frequency percentages based on https://en.wikipedia.org/wiki/Letter_frequency
func buildFrequencyMap() {
	freqMap['a']  = 8.167
	freqMap['b'] = 1.492
	freqMap['c'] = 2.782
	freqMap['d'] = 4.253
	freqMap['e'] = 12.702
	freqMap['f'] = 2.228
	freqMap['g'] = 2.015
	freqMap['h'] = 6.094
	freqMap['i'] = 6.966
	freqMap['j'] = 0.153
	freqMap['k'] = 0.772
	freqMap['l'] = 4.025
	freqMap['m'] = 2.406
	freqMap['n'] = 6.749
	freqMap['o'] = 7.507
	freqMap['p'] = 1.929
	freqMap['q'] = 0.095
	freqMap['r'] = 5.987
	freqMap['s'] = 6.327
	freqMap['t'] = 9.056
	freqMap['u'] = 2.758
	freqMap['v'] = 0.978
	freqMap['w'] = 2.360
	freqMap['x'] = 0.150
	freqMap['y'] = 1.974
	freqMap['z'] = 0.074
}
