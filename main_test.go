package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleDoCat() {
	// ファイルから読む場合
	doCat("test/test00")
	// Output:
	// test01
	// test02
	// test03
}

func ExampleDoCat2() {
	// 標準入力から読む場合
	content := []byte("test1234")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = tmpfile

	doCat()

	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// test1234

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
