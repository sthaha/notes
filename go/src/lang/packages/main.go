package main

import (
	"fmt"
	"lang/packages/shouter"
	"lang/packages/simplifier"
)

func main() {
	fmt.Println(shouter.Shout("Hello world"), " == HELLO WORLD")
	fmt.Println(simplifier.Simplify("Hello, world!"), " == hello world")
}
