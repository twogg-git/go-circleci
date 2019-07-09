package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseEmptyString(t *testing.T) {
	s := ""
	e := ""
	r := Flip(s)
	assert.Equal(t, e, r)
}

func TestReverseSingleCharString(t *testing.T) {
	s := "a"
	e := "a"
	r := Flip(s)
	assert.Equal(t, e, r)
}

func TestReverseTwoCharString(t *testing.T) {
	s := "ab"
	e := "ba"
	r := Flip(s)
	assert.Equal(t, e, r)
}

func TestReverseRepeatedCharsString(t *testing.T) {
	s := "aaa"
	e := "aaa"
	r := Flip(s)
	assert.Equal(t, e, r)
}

func TestReverseLongerString(t *testing.T) {
	s := "longer and complex words"
	e := "sdrow xelpmoc dna regnol"
	r := Flip(s)
	assert.Equal(t, e, r)
}
