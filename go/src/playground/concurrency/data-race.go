package main

import (
	"fmt"
	"os"
	"time"
)

var count = 0

func inc() {
	if count == 0 {
		count++
	}
}

func main() {
	go inc()
	go inc()

	time.Sleep(2 * time.Millisecond)
	fmt.Println("count:", count)

	if count != 1 {
		os.Exit(1)
	}
}
