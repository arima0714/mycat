package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args     = kingpin.Arg("filenames", "filenames").Strings()
	showends = kingpin.Flag("show-ends", "show \"$\" end of line").Short('E').Bool()
)

func doCat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	fileBuf := bufio.NewScanner(file)
	for i := 1; fileBuf.Scan(); i++ {
		line := fileBuf.Text()
		if *showends {
			line = line + "$"
		}
		fmt.Println(line)
	}
}

func main() {
	kingpin.Parse()
	filenames := *args

	if len(filenames) == 0 {
		for {
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			line := input.Text()
			if *showends {
				line = line + "$"
			}
			fmt.Println(line)
		}
	} else {
		for i := 0; i < len(filenames); i++ {
			filename := filenames[i]
			doCat(filename)
		}
	}
}
