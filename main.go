package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jzandbergen/promfmt/parser"
)

func main() {
	// we read from stdin
	reader := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)

	var s strings.Builder

	for {
		// read max CAPacity of buffer
		n, err := reader.Read(buf[:cap(buf)])
		// if data is smaller than buffer, trim data.
		buf = buf[:n]

		// no more data to read?
		if n == 0 {
			if err == nil {
				// nothing read, try again...
				continue
			}
			// error handing
			if err == io.EOF {
				break
			}
			fmt.Println(err)
			os.Exit(1)
		}

		// process buffer
		s.Write(buf)
	}

	// Thank you so much peoples of the prometheus for already providing
	// a pretty printer. The only reason I stole your parser code is because
	// the maxCharactersPerLine was not exported and I want the pretty code
	// to break the line earlier than 100 chars.
	parser.MaxCharactersPerLine = 10

	x, err := parser.ParseExpr(s.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	y := parser.Prettify(x)
	fmt.Println(y)
}
