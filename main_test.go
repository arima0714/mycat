package main_test

import (
	"bytes"
	"testing"

	"github.com/arima0714/mycat"
	"github.com/stretchr/testify/assert"
)

func resetShowends() {
	*main.ExpShowEnds = false
}

func ExampleExpDoCat() {
	defer resetShowends()

	var f bytes.Buffer
	f.Write([]byte("test1\ntest2\ntest3\n"))
	main.ExpDoCat(&f)

	*main.ExpShowEnds = true
	f.Write([]byte("あいう１\nあいう２\nあいう３"))
	main.ExpDoCat(&f)

	// Output:
	// test1
	// test2
	// test3
	// あいう１$
	// あいう２$
	// あいう３
}

func TestCheckEOL(t *testing.T) {
	defer resetShowends()

	*main.ExpShowEnds = false
	assert.Equal(t, main.ExpIsEOL('a'), false)

	*main.ExpShowEnds = false
	assert.Equal(t, main.ExpIsEOL('\n'), false)

	*main.ExpShowEnds = true
	assert.Equal(t, main.ExpIsEOL('a'), false)

	*main.ExpShowEnds = true
	assert.Equal(t, main.ExpIsEOL('\n'), true)
}
