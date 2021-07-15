package main

import (
	"flag"
	"fmt"
	"os"
)

func exists(name string) bool {
	_, err := os.Stat(name)
	return os.IsNotExist(err)
}

func main() {
	flag.Parse()
	args := flag.Args()
	// ファイルが存在しなければ「cat: <ファイル名>: No such file or directory」と表示
	for argIndex := 0; argIndex < len(args); argIndex++ {
		fileName := args[argIndex]
		if exists(fileName) {
			errString := fmt.Sprintf("cat: %s: No such file or directory\n", fileName)
			fmt.Fprint(os.Stdout, errString)
			continue
		} else {
			fmt.Println(fileName)
		}
		// ファイルを開く
		file, _ := os.Open(fileName)
		defer file.Close()
		// ファイル内のテキストを改行ごとに表示
	}
}
