package Crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"server/Logging"
)

func Encrypt(data []byte) (string, bool) {
	/*
		Encrypting the data with the provided key
		Parameters - data (to be encrypted), key (which will be used to encrypt the data)
		Returns - cipherText (encrypted data), status
	*/

	// Creating a AES cipher block
	block, _ := aes.NewCipher([]byte(HashMD5(Key)))

	// Creating a Galois/Counter Mode cipher block using the AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		Logging.ERROR.Println("Encryption failed, ", err)
		return "", false
	}
	// Creating Nonce  - a random data
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		Logging.ERROR.Println("Encryption failed, ", err)
		return "", false
	}

	// Encrypting and returning the data and the status
	return hex.EncodeToString(gcm.Seal(nonce, nonce, data, nil)), true
}

func Decrypt(data []byte) (string, bool) {
	/*
		Decrypting the data with the provided key
		Parameters - data (to be decrypted), key (which will be used to decrypt the data)
		Returns - plaintext (decrypted data), status
	*/

	// Creating a AES cipher block
	block, err := aes.NewCipher([]byte(HashMD5(Key)))
	if err != nil {
		Logging.ERROR.Println("Decryption failed, ", err)
		return "", false
	}
	// Creating a Galois/Counter Mode cipher block using the AES block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		Logging.ERROR.Println("Decryption failed, ", err)
		return "", false
	}
	// Getting the nonce size to extract nonce and cipher text from the input given data
	nonceSize := gcm.NonceSize()
	// Decode the input string to []byte
	data, err = hex.DecodeString(string(data))
	if err != nil {
		Logging.ERROR.Println("Decryption failed, ", err)
		return "", false
	}
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	// Decrypting the given cipher text with the key
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {
		Logging.ERROR.Println("Decryption failed, ", err)
		return "", false
	}
	// Decrypting and returning the data and the status
	return string(plaintext), true
}
