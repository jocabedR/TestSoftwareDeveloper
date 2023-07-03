package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateDate(t *testing.T) {

	validationCase1, _ := validateDate("45/13/2002")
	validationCase2, dateCase2 := validateDate("03/02/2001")
	expectDateCase2 := time.Date(2001, time.February, 3, 0, 0, 0, 0, time.UTC)
	validationCase3, _ := validateDate("02-04-2010")

	assert.Equal(t, validationCase1, false)
	assert.Equal(t, validationCase2, true)
	assert.Equal(t, dateCase2, expectDateCase2)
	assert.Equal(t, validationCase3, false)
}
