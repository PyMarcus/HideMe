package encrypt

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key = []byte("F1D2D2F924E986AC86FDF7B36C94BCD1")
)

func TestEncryptFile(t *testing.T) {
	input := "input.txt"
	err := EncryptFile(key, input)
	assert.NoError(t, err)
}

func TestCreateCipher(t *testing.T){
	c, err := createCipher(key)
	
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestCreateGCM(t *testing.T){
	block, _ := createCipher(key)

	c, err := createGCM(block)
	
	assert.NoError(t, err)
	assert.NotNil(t, c)
}

func TestReadFile(t *testing.T){
	input := "input.txt"

	inFile, _ := os.OpenFile(input, os.O_RDWR, 0600)

	_, err := readFile(inFile)
	assert.NoError(t, err)
}