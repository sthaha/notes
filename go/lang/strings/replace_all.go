package main

import (
	"fmt"
	"strings"
)

type x int

func (s x) String() string {
	return "oo"
}

func main() {

	text := "foobar"
	find := x(0)
	replacement := "00"

	replaced := strings.Replace(text, find, replacement, -1)
	fmt.Println(replaced)
}
