package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareStrings(t *testing.T) {
	equal := CompareStrings("правильный ответ", "правильный ответ")
	assert.Equal(t, equal, true)

	equal = CompareStrings("Правильный Ответ", "правильный ответ")
	assert.Equal(t, equal, true)

	equal = CompareStrings("!правильный ответ?", "правильный ответ")
	assert.Equal(t, equal, true)

	equal = CompareStrings("правильный вопрос", "правильный ответ")
	assert.Equal(t, equal, false)
}
