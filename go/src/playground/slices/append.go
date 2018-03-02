package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	a = append(a, b...)

	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
