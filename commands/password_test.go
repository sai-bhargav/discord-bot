package commands

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordGenarator(t *testing.T) {
	assert.Equal(t, "", generator("pwd 0"))
	assert.Equal(t, 12, len(generator("pwd 12")))
	assert.Equal(t, 123, len(generator("pwd 123")))

	assert.Equal(t, 123, len(generator("pwd 123456789876546788")))
	assert.Equal(t, "", generator("pwd 000456789876546788"))

	assert.Equal(t, "Please check the input", generator("pwd asdc"))
	assert.Equal(t, "Please check the input", generator("pws 23"))
}

func Regression(t *testing.T) {
	assert.True(t, true, "True is true!")
}
