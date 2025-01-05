package main

import (
	"os"

	"github.com/senicko/writing-an-interpreter-in-go/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
