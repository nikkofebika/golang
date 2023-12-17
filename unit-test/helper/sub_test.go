package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Test1", func(t *testing.T) {
		name := "nikko"
		result := HelloWorld(name)

		assert.Equal(t, "Hello "+name, result, "Should be Hello "+name)
	})
	t.Run("Test2", func(t *testing.T) {
		name := "nikko"
		result := HelloWorld(name)

		assert.Equal(t, "Hello "+name, result, "Should be Hello "+name)
	})
}
