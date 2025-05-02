package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

func EncryptFile(key []byte, inputFile string) error {
	inFile, err := os.OpenFile(inputFile, os.O_RDWR, 0600)
	if err != nil {
		return err
	}
	defer inFile.Close()

	plainData, err := readFile(inFile)
	if err != nil {
		return err
	}

	cipherBlock, err := createCipher(key)
	if err != nil {
		return err
	}

	gcm, err := createGCM(cipherBlock)
	if err != nil {
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return err
	}

	cipherData := gcm.Seal(nil, nonce, plainData, nil)

	if _, err := inFile.Seek(0, 0); err != nil {
		return err
	}

	err = writeFile(inFile, nonce, cipherData)
	if err != nil {
		return err
	}

	if err := inFile.Truncate(int64(len(nonce) + len(cipherData))); err != nil {
		return err
	}
	
	return nil
}

func writeFile(inFile *os.File, nonce []byte, dataEncrypted []byte) error {
	if _, err := inFile.WriteAt(nonce, 0); err != nil {
		return err
	}

	if _, err := inFile.WriteAt(dataEncrypted, int64(len(nonce))); err != nil {
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

func readFile(file *os.File) ([]byte, error) {
	plainData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return plainData, nil
}
