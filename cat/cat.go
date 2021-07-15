package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()
	// ファイルが存在しなければ「cat: <ファイル名>: No such file or directory」と表示
	for argIndex := 0; argIndex < len(args); argIndex++ {
		fmt.Println(args[argIndex])
	}
}
