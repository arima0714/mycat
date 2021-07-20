package main

import (
	"bufio"
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
		}
		// ファイルを開く
		file, _ := os.Open(fileName)
		defer file.Close()
		// ファイル内のテキストを改行ごとに表示
		fileBuf := bufio.NewScanner(file)
		for i := 1; fileBuf.Scan(); i++ {
			line := fileBuf.Text()
			fmt.Println(line)
		}
	}
	// if len(args) == 0 {
	// 	for ;; {
	// 		var input string
	// 		fmt.Scanln(input)
	// 		fmt.Println(input)
	// 	}
	// }
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
