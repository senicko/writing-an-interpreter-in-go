package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/senicko/writing-an-interpreter-in-go/evaluator"
	"github.com/senicko/writing-an-interpreter-in-go/lexer"
	"github.com/senicko/writing-an-interpreter-in-go/parser"
)

const PROMPT = ">> "

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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error!")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
