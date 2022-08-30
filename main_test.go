package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doCat(t *testing.T) {

}

func Test_processLine(t *testing.T) {
	var line *string
	*line = "line"
	var expected string = "line$"
	var actually *string
	processLine(line)
	actually = line
	assert.Equal(t, expected, actually)
}
