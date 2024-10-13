package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args

	if(args[1] == "enc_hex") {fmt.Print(enc_hex(args[2]))}
	if(args[1] == "dec_hex") {fmt.Print(dec_hex(args[2]))}

	if(args[1] == "enc") {
		hashedKey := sha256.Sum256([]byte(args[2]))
		aesKey := hashedKey[:]
		encryptedText, err := encrypt(args[3], string(aesKey))
		if err != nil {
			fmt.Print("error:21:", err)
			return
		}

		fmt.Print("sucess:" + encryptedText)
	}

	if args[1] == "dec" {
		hashedKey := sha256.Sum256([]byte(args[2]))
		aesKey := hashedKey[:]
		decryptedText, err := decrypt(args[3], string(aesKey))
		if err != nil {
			fmt.Println("error:21:", err)
			return
		}

		fmt.Print("sucess:" + decryptedText)
	}

	if (args[1] == "hash") {
		fmt.Print(sha256.Sum256([]byte(args[2])))
	}
}

func enc_hex(s string) string {
	hexString := ""
	for _, b := range []byte(s) {
		hexString += fmt.Sprintf("%02X", b)
	}
	return hexString
}

func dec_hex(hexStr string) (string) {
    bytes, err := hex.DecodeString(hexStr)
    if err != nil {
        return ""
    }
    return string(bytes)
}

// encrypt encrypts the plain text using the given key.
func encrypt(plainText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherText := aesGCM.Seal(nonce, nonce, []byte(plainText), nil)
	return hex.EncodeToString(cipherText), nil
}

// decrypt decrypts the cipher text using the given key.
func decrypt(cipherText, key string) (string, error) {
	data, err := hex.DecodeString(cipherText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	nonce, cipherTextBytes := data[:nonceSize], data[nonceSize:]

	plainText, err := aesGCM.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plainText), nil
}