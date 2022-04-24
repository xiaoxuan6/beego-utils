package bcrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	hash := Password("123456")
	re := ValidPassword(hash, "123456")

	assert.Equal(t, true, re)
}
