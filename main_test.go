package main_test

import (
	"bytes"

	"github.com/arima0714/mycat"
)

func ExampleDoCat() {
	var stdin bytes.Buffer
	stdin.Write([]byte("test1234"))

	main.DoCat(&stdin)

	// Output:
	// test1234
}

func ExampleDoCat2() {
	var stdin bytes.Buffer
	*main.Showends = true
	stdin.Write([]byte("test1234"))

	main.DoCat(&stdin)

	// Output:
	// test1234$
}
