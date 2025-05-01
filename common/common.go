package common

import (
	"crypto/pbkdf2"
	"crypto/sha256"
	"log"
)

const (
	INTERACTIONS int = 100000
	BYTES        int = 32
)

func GenerateKeyFromPassword(password string) []byte {
	passwordByte := []byte(password)
	salt := passwordByte

	key, err := pbkdf2.Key(sha256.New, password, salt, INTERACTIONS, BYTES)

	if err != nil{
		log.Fatal("[-] Fail to generate key.")
	}

	return key
}
