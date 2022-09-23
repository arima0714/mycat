package main_test

import (
	"bytes"

	"github.com/arima0714/mycat"
)

func ExampleDoCat() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("test1"))

	main.ExpDoCat(&f)

	// Output:
	// test1
}

func ExampleDoCat2() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("test1\ntest2\ntest3\n"))

	main.ExpDoCat(&f)

	// Output:
	// test1$
	// test2$
	// test3$
}

func ExampleDoCat3() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("test1\ntest2\ntest3"))
	main.ExpDoCat(&f)

	// Output:
	// test1$
	// test2$
	// test3
}

func ExampleDoCat4() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("あああ１\nあああ２\nあああ３"))
	main.ExpDoCat(&f)

	// Output:
	// あああ１$
	// あああ２$
	// あああ３
}
