package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doCat(t *testing.T) {

}

func Test_processLine(t *testing.T) {
	*showends = true
	line := "line"
	var expected string = "line$"
	var actual string
	processLine(&line)
	actual = line
	assert.Equal(t, expected, actual)

	*showends = false
	line = "line"
	expected = "line"
	processLine(&line)
	actual = line
	assert.Equal(t, expected, actual)
}
