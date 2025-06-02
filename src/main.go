package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lines, words, characters int
	insideWord := false

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Fprintf(os.Stderr, "%s: cannot read stdin: %v\n", os.Args[0], err)
			os.Exit(74)
		}

		characters++

		if r == '\n' {
			lines++
		}

		if unicode.IsSpace(r) {
			insideWord = false
		} else if !insideWord {
			insideWord = true
			words++
		}
	}

	fmt.Printf("%8d%8d%8d\n", lines, words, characters)
}
