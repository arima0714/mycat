package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args     = kingpin.Arg("filenames", "filenames").Strings()
	showends = kingpin.Flag("show-ends", "show \"$\" end of line").Short('E').Bool()
)

func doCat(filename ...string) {
	pipedFile, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if pipedFile.Size() > 0 {
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		lines := strings.Split(string(b), "\n")
		for i := 0; i < len(lines); i++ {
			line := lines[i]
			processLine(&line)
			fmt.Println(line)
		}
	} else if len(filename) == 0 {
		for {
			input := bufio.NewScanner(os.Stdin)
			input.Scan()
			line := input.Text()
			processLine(&line)
			fmt.Println(line)
		}
	} else {
		file, err := os.Open(filename[0])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		fileBuf := bufio.NewScanner(file)
		for i := 1; fileBuf.Scan(); i++ {
			line := fileBuf.Text()
			processLine(&line)
			fmt.Println(line)
		}
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
