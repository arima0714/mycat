package main

import (
	"bufio"
	"fmt"
	"strings"
)

var source = `1行目
2行目
3行目
`

func main() {
	reader := bufio.NewReader(strings.NewReader(source))
	for {
		line, _ := reader.ReadString('\n')
		fmt.Printf("%#v\n", line)
		break
	}
}
