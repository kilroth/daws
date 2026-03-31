package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"

	"github.com/denisbrodbeck/machineid"
)

//var masterKey = []byte("my-very-strong-master-key") // In production, use a secure key management solution

func getMasterKey() ([]byte, error) {
	id, err := machineid.ID()
	if err != nil {
		id = "daws-app-fallback-id"
	}
	hash := sha256.Sum256([]byte(id + "a-unique-salt"))
	return hash[:], nil
}

func MaskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:4] + "****************"
}

func isMasked(key string) bool {
	if len(key) <= 8 {
		return key == "****"
	}
	return len(key) > 8 && key[4:8] == "****"
}

func Encrypt(plainText string) (string, error) {
	if len(plainText) > 50 {
		return plainText, nil
	}
	if plainText == "" {
		return "", Logger.Error("Cannot encrypt empty string")
	}
	masterKey, err := getMasterKey()
	if err != nil {
		return "", Logger.Error("Failed to get master key: %v", err)
	}
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", Logger.Error("Failed to create cipher block: %v", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", Logger.Error("Failed to create GCM: %v", err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", Logger.Error("Failed to generate nonce: %v", err)
	}
	cipherText := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func Decrypt(cipherText string) (string, error) {
	if cipherText == "" {
		return "", Logger.Error("Cannot decrypt empty string")
	}
	masterKey, err := getMasterKey()
	if err != nil {
		return "", Logger.Error("Failed to get master key: %v", err)
	}
	data, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		return "", Logger.Error("Failed to decode cipher text: %v", err)
	}
	block, err := aes.NewCipher(masterKey)
	if err != nil {
		return "", Logger.Error("Failed to create cipher block: %v", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", Logger.Error("Failed to create GCM: %v", err)
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", Logger.Error("Cipher text too short")
	}
	nonce, cipherTextBytes := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherTextBytes, nil)
	if err != nil {
		return "", Logger.Error("Failed to decrypt: %v", err)
	}
	return string(plainText), nil
}
