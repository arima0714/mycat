package main_test

import (
	"bytes"
)

func ExampleDoCat() {
	var stdin bytes.Buffer
	stdin.Write([]byte("test1234"))

	main.ExpDoCat(&stdin)

	// Output:
	// test1234
}

func ExampleDoCat2() {
	var stdin bytes.Buffer
	*main.ExpShowends = true
	stdin.Write([]byte("test1234\n1234test\n1t2e3s4t\n"))

	main.ExpDoCat(&stdin)

	// Output:
	// test1234$
	// 1234test$
	// 1t2e3s4t$
	*main.ExpShowends = false
}

func ExampleDoCat3() {
	var stdin bytes.Buffer
	*main.ExpShowends = true
	stdin.Write([]byte("test01\ntest02\ntest03"))
	main.ExpDoCat(&stdin)

	// Output:
	// test01$
	// test02$
	// test03
}
