package lib

import "reflect"

// the number of matches we consider to be to unlikely to be coincidence
var MATCH_THRESHOLD int = 2

// CheckECB compares each 16 byte block in the input to every other 16 byte block in
// the input, checking to see if there are matching blocks
func CheckECB(bytes []byte) bool {
	numMatches := 0
	for i := 0; i < (len(bytes) / 16) - 1; i++ {
		block := bytes[(i * 16):((i + 1) * 16)]
		for j := i + 1; j < (len(bytes) / 16) - 2; j++ {
			otherBlock := bytes[(j * 16):((j + 1) * 16)]
			if (reflect.DeepEqual(block, otherBlock)) {
				numMatches++
			}
		}
	}
	if numMatches >= MATCH_THRESHOLD {
		return true
	}
	return false
}
