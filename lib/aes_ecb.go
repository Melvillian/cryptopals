package lib

import (
	"crypto/cipher"
)

// AESECBModeDecrypt decrypts some cipherText using a Block created by crypt/aes.NewCipher
func AESECBModeDecrypt(cipherText []byte, block cipher.Block) []byte {
	bs := block.BlockSize()
	if len(cipherText) % bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	var plainText []byte
	for j := 0; j < len(cipherText) / bs; j++ {
		decryptedBlock := make([]byte, bs)
		block.Decrypt(decryptedBlock, cipherText)
		plainText = append(plainText, decryptedBlock...)
		cipherText = cipherText[bs:]
	}
	return plainText
}

// AESECBModeEncrypt encrypts some plaintext using a Block created by crypt/aes.NewCipher
func AESECBModeEncrypt(plaintext []byte, block cipher.Block) []byte {
	bs := block.BlockSize()
	if len(plaintext) % bs != 0 {
		panic("Need a multiple of the blocksize")
	}

	var cipherText []byte
	for j := 0; j < len(cipherText) / bs; j++ {
		encryptedBlock := make([]byte, bs)
		block.Decrypt(encryptedBlock, cipherText)
		cipherText = append(cipherText, encryptedBlock...)
		plaintext = plaintext[bs:]
	}
	return cipherText
}