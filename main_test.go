package main_test

import (
	"bytes"
	"github.com/arima0714/mycat"
	"log"
	"os"
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
	stdin.Write([]byte("test1234\n1234test\n1t2e3s4t\n"))

	main.DoCat(&stdin)

	// Output:
	// test1234$
	// 1234test$
	// 1t2e3s4t$
	*main.Showends = false
}

func ExampleDoCat3() {
	*main.Showends = true
	f, err := os.Open("./test/test00")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	main.DoCat(f)

	// Output:
	// test01$
	// test02$
	// test03
}
