package main_test

import (
	"bytes"
	"testing"

	"github.com/arima0714/mycat"
	"github.com/stretchr/testify/assert"
)

func resetShowends() {
	*main.ExpShowends = false
}

func ExampleExpDoCat() {
	defer resetShowends()
	
	var f bytes.Buffer
	f.Write([]byte("test1\ntest2\ntest3\n"))
	main.ExpDoCat(&f)

	*main.ExpShowends = true
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

	*main.ExpShowends = false
	assert.Equal(t, main.ExpCheckEOL('a'), false)

	*main.ExpShowends = false
	assert.Equal(t, main.ExpCheckEOL('\n'), false)
	
	*main.ExpShowends = true
	assert.Equal(t, main.ExpCheckEOL('a'), false)
	
	*main.ExpShowends = true
	assert.Equal(t, main.ExpCheckEOL('\n'), true)
}
