package main_test

import (
	"bytes"

	"github.com/arima0714/mycat"
)

func ExampleDoCat() {
	var f bytes.Buffer
	f.Write([]byte("test1234"))

	main.ExpDoCat(&f)

	// Output:
	// test1234
}

func ExampleDoCat2() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("test1234\n1234test\n1t2e3s4t\n"))

	main.ExpDoCat(&f)

	// Output:
	// test1234$
	// 1234test$
	// 1t2e3s4t$
	*main.ExpShowends = false
}

func ExampleDoCat3() {
	var f bytes.Buffer
	*main.ExpShowends = true
	f.Write([]byte("test01\ntest02\ntest03"))
	main.ExpDoCat(&f)

	// Output:
	// test01$
	// test02$
	// test03
}
