package lib

// PadWithPKCS7 pads the given byte slice to the given blockSize using the PKCS#7 padding scheme,
// returning the padded input slice
func PadWithPKCS7(data []byte, blockSize int) []byte {
	paddedData := data

	padLen := blockSize - (len(data) % blockSize)
	for i := 0; i < padLen; i++ {
		paddedData = append(paddedData, byte(padLen))
	}
	return paddedData
}
