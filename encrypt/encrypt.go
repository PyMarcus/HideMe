package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func EncryptFile(key []byte, inputFile, outputFile string) error{
	inFile, err := os.Open(inputFile)

	if err != nil{
		log.Println("[-] Fail to open file " + inputFile)
		return err 
	}

	defer inFile.Close()

	outFile, err := os.Create(outputFile)

	if err != nil{
		log.Println("[-] Fail to create file " + outputFile)
		return err 
	}

	defer outFile.Close()

	return nil
}

func createCipher(key []byte) (cipher.Block, error){
	block, err := aes.NewCipher(key)

	if err != nil{
		log.Println("[-] Fail to create cypher")
		return nil, err
	}

	return block, nil
}