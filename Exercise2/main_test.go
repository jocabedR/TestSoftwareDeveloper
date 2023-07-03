package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNumbers(t *testing.T) {

	assert.Equal(t, 8, countNumbers("jj@dud324khsa123h12"))
	assert.Equal(t, 7, countNumbers("jj@dud324khsa123h1"))
	assert.Equal(t, 0, countNumbers("jj@dud"))
	assert.Equal(t, 0, countNumbers(""))
	assert.Equal(t, 10, countNumbers("1111111111"))
	assert.NotEqual(t, 7, countNumbers("jj@dud324khsa123h12"))
}
