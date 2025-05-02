package decrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"io"
	"os"
)

func DecryptFile(key []byte, inputFile string) error {
	inFile, err := os.OpenFile(inputFile, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer inFile.Close()

	cipherBlock, err := createCipher(key)
	if err != nil {
		return err
	}

	gcm, err := createGCM(cipherBlock)
	if err != nil {
		return err
	}

	encryptedData, err := io.ReadAll(inFile)
	if err != nil {
		return err
	}

	nonceSize := gcm.NonceSize()
	if len(encryptedData) < nonceSize {
		return io.ErrUnexpectedEOF
	}

	nonce := encryptedData[:nonceSize]
	cipherData := encryptedData[nonceSize:]

	plainData, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return err
	}

	if _, err := inFile.Seek(0, 0); err != nil {
		return err
	}

	if _, err := inFile.Write(plainData); err != nil {
		return err
	}

	if err := inFile.Truncate(int64(len(plainData))); err != nil {
		return err
	}

	return nil
}

func createCipher(key []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func createGCM(c cipher.Block) (cipher.AEAD, error) {
	aesGCM, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	return aesGCM, nil
}
