package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func test_doCat(t *testing.T){ 

}

func test_processLine(t *testing.T) {
	*bool showends = true
	*string line = "line"
	expected = "line$"
	actually = processLine(line)
	assert (t, expected, actually)
}

