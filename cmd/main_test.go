package main

import (
	"os"
	"testing"

	"github.com/PyMarcus/hideme/decrypt"
	"github.com/PyMarcus/hideme/encrypt"
	"github.com/stretchr/testify/assert"
)

var (
	key = []byte("F1D2D2F924E986AC86FDF7B36C94BCD1")
)

func TestEncryptAndDecryptIntegration(t *testing.T) {
	inputPath := "input_test_file.txt"
	originalContent := []byte("Hello, world!")

	err := os.WriteFile(inputPath, originalContent, 0600)
	assert.NoError(t, err)

	err = encrypt.EncryptFile(key, inputPath)
	assert.NoError(t, err)

	err = decrypt.DecryptFile(key, inputPath)
	assert.NoError(t, err)

	data, err := os.ReadFile(inputPath)
	assert.NoError(t, err)
	assert.Equal(t, originalContent, data)
}