package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/senicko/writing-an-interpreter-in-go/lexer"
	"github.com/senicko/writing-an-interpreter-in-go/token"
)

const PROMPT = ">>"

// Start starts the REPL with the given io.Reader and io.Writer.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
