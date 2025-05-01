package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryptFile(t *testing.T) {
	key := []byte("F1D2D2F924E986AC86FDF7B36C94BCDF32BEEC15F1F4A59A1D1B8375C6A75851")
	input := "input.txt"
	output := "output.txt"
	err := EncryptFile(key, input, output)
	assert.NoError(t, err)
}
