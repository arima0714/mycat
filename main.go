package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	args     = kingpin.Arg("filenames", "filenames").Strings()
	showends = kingpin.Flag("show-ends", "show \"$\" end of line").Short('E').Bool()
)

func doCat(f io.Reader) {
	buf := bufio.NewScanner(f)
	for buf.Scan() {
		line := buf.Text()
		if *showends {
			line += "$"
		}
		fmt.Println(line)
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
