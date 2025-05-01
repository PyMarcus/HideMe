package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key = []byte("F1D2D2F924E986AC86FDF7B36C94BCDW")
)

func TestEncryptFile(t *testing.T) {
	input := "input.txt"
	output := "output.txt"
	err := EncryptFile(key, input, output)
	assert.NoError(t, err)
}

func TestCreateCipher(t *testing.T){
	c, err := createCipher(key)
	
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
