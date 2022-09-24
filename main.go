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

func checkRune(r rune) bool {
	return *showends && r == '\n'
}

func doCat(f io.Reader) {
	r := bufio.NewReader(f)
	for {
		r, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		if checkRune(r) {
			fmt.Print("$")
		}
		fmt.Print(string(r))
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
