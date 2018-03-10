package main

import (
	"fmt"
	"playground/packages/shouter"
	"playground/packages/simplifier"
)

func main() {
	fmt.Println(shouter.Shout("Hello world"), " == HELLO WORLD")
	fmt.Println(simplifier.Simplify("Hello, world!"), " == hello world")
}
