package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)


func DecryptFile(key []byte, filePath string) error {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0600)
	
	if err != nil {
		return err
	}
	defer file.Close()

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = file.Read(nonce)
	if err != nil {
		return err
	}

	cipherData, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	plainData, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, plainData, 0600)
	return err
}