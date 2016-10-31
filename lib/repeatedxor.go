package lib

// RepeatedXor iterates through plaintext and encrypts each byte with the next byte in the key,
// returning to the first byte in the key once the last byte of the key has been used
func RepeatedXor(plaintext []byte, key []byte) (ciphertext []byte) {
	ciphertext = make([]byte, len(plaintext))
	for i, b := range plaintext {
		ciphertext[i] = b ^ key[i % len(key)]
	}

	return
}
