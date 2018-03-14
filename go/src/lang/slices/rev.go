package main

import "fmt"

func reverseInts(src []int) []int {
	dst := make([]int, len(src))

	for s, d := len(src)-1, 0; s >= 0; s, d = s-1, d+1 {
		dst[d] = src[s]
	}

	return dst
}

func main() {
	orginal = []int{4, 2, 3, 2, 7}
	fmt.Println("Reverse a slice", original, " -> ", reverseInts(orignal))
}
