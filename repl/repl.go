package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/senicko/writing-an-interpreter-in-go/compiler"
	"github.com/senicko/writing-an-interpreter-in-go/lexer"
	"github.com/senicko/writing-an-interpreter-in-go/parser"
	"github.com/senicko/writing-an-interpreter-in-go/vm"
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

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Whoops! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Whoops! Executing bytecode failed:\n %s\n", err)
			continue
		}

		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Error!")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
