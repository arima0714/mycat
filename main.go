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

func doCat(filename ...string) {
	var buf *bufio.Scanner
	if len(filename) == 0 {
		buf = bufio.NewScanner(os.Stdin)
	} else {
		file, err := os.Open(filename[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		buf = bufio.NewScanner(file)

	}
	for buf.Scan(){
		line := buf.Text()
		processLine(&line)
		fmt.Println(line)
	}
}

func processLine(line *string) {
	if *showends {
		*line = *line + "$"
	}
}

func main() {
	kingpin.Parse()
	filenames := *args

	if len(filenames) == 0 {
		doCat()
	} else {
		for i := 0; i < len(filenames); i++ {
			filename := filenames[i]
			doCat(filename)
		}
	}
}
