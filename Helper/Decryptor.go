package Helper

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"
)

// PKCS7PaddingRemoval removes the PKCS7 padding from the decrypted data
func PKCS7PaddingRemoval(data []byte) ([]byte, error) {
	padding := data[len(data)-1]
	if int(padding) > len(data) {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-int(padding)], nil
}

// DecryptData decrypts the given base64-encoded encrypted data
func DecryptData(encryptedData string) ([]byte, error) {
	// Decode the base64-encoded encrypted data
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	// The IV is the first 16 bytes of the ciphertext
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]

	// The actual encrypted data starts after the IV
	ciphertext = ciphertext[aes.BlockSize:]

	// Retrieve the AES key from an environment variable
	aesKey := []byte(os.Getenv("AES_KEY"))
	if len(aesKey) > 32 {
		aesKey = aesKey[:32] // Use only the first 32 bytes (for AES-256)
	}

	// Generate the AES block cipher
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}

	// Decrypt the data using AES in CBC mode
	mode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(ciphertext))
	mode.CryptBlocks(decrypted, ciphertext)

	// Remove the padding from the decrypted data
	decryptedData, err := PKCS7PaddingRemoval(decrypted)
	if err != nil {
		return nil, err
	}

	return decryptedData, nil
}
