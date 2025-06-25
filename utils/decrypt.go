package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
	"os"
	"strings"
)

func DecryptFile(path, password string, overwrite bool) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		if _, err := os.Stat(path + ".locked"); err == nil {
			path = path + ".locked"
		} else {
			fmt.Println("Error: file not found:", path)
			os.Exit(1)
		}
	}

	keyBytes := passwordToKey(password)

	ciphertext, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	outFileName := path
	if strings.HasSuffix(path, ".locked") {
		outFileName = path[:len(path)-len(".locked")]
	}

	if overwrite {
		err = os.Remove(path)
		if err != nil {
			fmt.Println("Error: Could not delete encrypted file:", err)
		}
	}

	return os.WriteFile(outFileName, plaintext, 0644)

}
