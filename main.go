package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args     = kingpin.Arg("filenames", "filenames").Strings()
	showends = kingpin.Flag("show-ends", "show \"$\" end of line").Short('E').Bool()
)

func doCat(f io.Reader) {
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		line = strings.Replace(line, "\n", "", -1)
		fmt.Print(line)
		if err == io.EOF {
			break
		} else {
			endOfLine := ""
			if *showends {
				endOfLine += "$"
			}
			fmt.Println(endOfLine)
		}
	}
}

func main() {
	kingpin.Parse()
	filenames := *args

	if len(filenames) == 0 {
		doCat(os.Stdin)
	} else {
		for _, file := range filenames {
			f, err := os.Open(file)
			if err != nil {
				log.Fatal(err)
			}
			defer f.Close()
			doCat(f)
		}
	}
}
