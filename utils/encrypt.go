package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func EncryptFile(path, name, password string, overwrite bool) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Println("Error: file not found:", path)
		os.Exit(1)
	}

	keyBytes := passwordToKey(password)

	if len(keyBytes) != 32 {
		return errors.New("Error: key is not 32-bit")
	}

	plaintext, err := os.ReadFile(path)
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	outPath := name

	if outPath == "" {
		outPath = path + ".locked"
	} else {
		if !filepath.IsAbs(outPath) {
			dir := filepath.Dir(path)
			outPath = filepath.Join(dir, outPath)
		}

		if filepath.Ext(outPath) == "" {
			originalExt := filepath.Ext(path)
			outPath += originalExt
		}
		outPath += ".locked"
	}

	if overwrite {
		err = os.Remove(path)
		if err != nil {
			fmt.Println("Error: Could not delete original file:", err)
		}
	}

	return os.WriteFile(outPath, ciphertext, 0644)
}
