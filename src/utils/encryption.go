package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/joho/godotenv"
)

var encryptionKey []byte

func init() {
	_ = godotenv.Load()
    key := os.Getenv("ENCRYPTION_SECRET_KEY")
	fmt.Println("ENCRYPTION_SECRET_KEY: ", os.Getenv("JWT_SECRET_KEY"))
    if key == "" {
        panic("ENCRYPTION_SECRET_KEY environment variable not set")
    }
    if len(key) != 32 { // AES-256 needs 32 bytes key
        panic("ENCRYPTION_SECRET_KEY must be 32 bytes long")
    }
    encryptionKey = []byte(key)
}

// Encrypt encrypts plain text string and returns base64 encoded string
func Encrypt(plainText string) (string, error) {
    block, err := aes.NewCipher(encryptionKey)
    if err != nil {
        return "", err
    }

    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]

    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], []byte(plainText))

    return base64.URLEncoding.EncodeToString(cipherText), nil
}

// Decrypt decrypts base64 encoded string
func Decrypt(encryptedText string) (string, error) {
    cipherText, err := base64.URLEncoding.DecodeString(encryptedText)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher(encryptionKey)
    if err != nil {
        return "", err
    }

    if len(cipherText) < aes.BlockSize {
        return "", errors.New("cipherText too short")
    }

    iv := cipherText[:aes.BlockSize]
    cipherText = cipherText[aes.BlockSize:]

    stream := cipher.NewCFBDecrypter(block, iv)
    stream.XORKeyStream(cipherText, cipherText)

    return string(cipherText), nil
}
