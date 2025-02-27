package main

import (
	"fmt"
	"os"

	"github.com/senicko/writing-an-interpreter-in-go/evaluator"
	"github.com/senicko/writing-an-interpreter-in-go/lexer"
	"github.com/senicko/writing-an-interpreter-in-go/object"
	"github.com/senicko/writing-an-interpreter-in-go/parser"
	"github.com/senicko/writing-an-interpreter-in-go/repl"
)

func main() {
	if len(os.Args) > 1 {
		source, err := os.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("Invalid monkey source file!")
			os.Exit(1)
		}

		execute(string(source))
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}

// TODO: move this logic somewhere else
func execute(source string) {
	l := lexer.New(source)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluator.Eval(program, env)
}
