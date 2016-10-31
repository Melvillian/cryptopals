package lib

import (
	"errors"
)


// HammingDist finds the number of different bits between two bytes slices. Both
// slices must be the same length, else an error will be returned
func HammingDist(b1 []byte, b2 []byte) (uint32, error) {
	if (len(b1) != len(b2)) {
		return 0, errors.New("arguments must have an equal number of bytes");
	}

	// xoring the bytes gives us a 0 everywhere the bits match and a 1 everywhere they differ
	xor := XorBytes(b1, b2)

	totalBits := uint32(0)
	for _, b := range xor {
		totalBits += numBits(b)
	}

	return totalBits, nil
}

// numBits returns the number of 1's in the given byte
func numBits(b byte) uint32 {
	totalBits := uint32(0)
	for i := uint8(0); i < 8; i++ {
		totalBits += uint32((b >> i) & 1)
	}
	return totalBits
}
