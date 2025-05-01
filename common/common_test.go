package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeyFromPassword(t *testing.T){
	testKey := GenerateKeyFromPassword("myPassword")
	assert.NotEmpty(t, testKey)
	assert.Equal(t, len(testKey), BYTES)
}