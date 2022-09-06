package helpers

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)


func TestGetEnv(t *testing.T) {
	assert.Equal(t, GetEnv("ABCDEFG","fallback"), "fallback")
	os.Setenv("ABCDEFG", "notFallBack")
	assert.Equal(t, GetEnv("ABCDEFG","fallback"), "notFallBack")
}

