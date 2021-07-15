package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	// ファイルが存在しなければ「cat: <ファイル名>: No such file or directory」と表示
	for argIndex := 0; argIndex < len(flag.Args()); argIndex++{
		fmt.Println(argIndex)
	}
}
