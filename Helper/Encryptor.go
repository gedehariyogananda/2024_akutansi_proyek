package Helper

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"os"
)

// PKCS7Padding is PKCS7 padding function
func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func EncryptData(data []byte) (string, error) {
	aesKey := []byte(os.Getenv("AES_KEY"))
	if len(aesKey) > 32 {
		aesKey = aesKey[:32]
	}

	// Generate the AES block cipher
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", err
	}

	// Pad the data to a multiple of the block size (16 bytes for AES)
	padData := PKCS7Padding(data, aes.BlockSize)

	// Create a ciphertext slice to hold the encrypted data (IV + padded data)
	ciphertext := make([]byte, aes.BlockSize+len(padData))

	// Generate a random initialization vector (IV)
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Encrypt the data using AES in CBC mode
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], padData)

	// Return the encrypted data in base64 encoding (to store in MongoDB)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
