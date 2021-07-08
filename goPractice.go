package main

import (
	"bytes"
	"fmt"
)

func main() {
	// Write()メソッドで書き込まれた内容をためておいて、後でまとめて結果を受け取れるbytes.Buffer
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}
