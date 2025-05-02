package decrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	key = []byte("F1D2D2F924E986AC86FDF7B36C94BCD1")
)

func TestDecryptFile(t *testing.T){
	err := DecryptFile(key, "../encrypt/input.txt")
	assert.NoError(t, err)
}