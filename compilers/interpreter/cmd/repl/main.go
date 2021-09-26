package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/sthaha/interpreter/lexer"
	"github.com/sthaha/interpreter/token"
)

func repl(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)

	for {
		fmt.Fprintf(w, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		input := scanner.Text()
		fmt.Printf("input: %s", input)
		lxr := lexer.New(input)

		fmt.Fprintf(w, "\n ---------------------------------------\n")
		for tkn := lxr.Next(); tkn.Type != token.EOF; tkn = lxr.Next() {
			fmt.Fprintf(w, "\n ... %v", tkn)
		}

		fmt.Fprintf(w, "\n ---------------------------------------\n")
	}
}
func main() {
	repl(os.Stdin, os.Stdout)
}
