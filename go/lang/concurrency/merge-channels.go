package main

import (
	"fmt"
	"math/rand"
	"time"
)

func nilChannel() <-chan int {
	return nil
}

func withValue() <-chan int {
	ch := make(chan int)

	time.AfterFunc(300*time.Millisecond, func() {
		defer close(ch)
		ch <- 42
	})
	return ch
}

func main() {
	// start by printing a nil channel

	c := nilChannel()
	v := withValue()

	fmt.Printf("Printing a nil channel: %+v \n", c)
	fmt.Printf("Printing a non-nil channel: %+v \n", v)

	// this ends in a deadlock; don't know why
	// fmt.Println("Read from a nil channel: %+v", <-c)

	// using select is safe and the operation is no-op
	select {
	case r, ok := <-c:
		fmt.Printf("This should never happen %+v, %b \n", r, ok)
	case r, ok := <-v:
		fmt.Printf("Using select between nil and non-nil does not deadlock: %+v, %b \n", r, ok)
	}

	// reading a closed channel returns 0 value

	fmt.Printf("Reading a closed channel gives Zero value: %+v \n", <-v)
	fmt.Printf("Reading a closed channel again: %+v \n", <-v)

	fmt.Printf("\n\nUsing the above to merge channels \n")

	chA := makeProducer([]int{0, 2, 4, 6, 8, 10, 12, 14})
	chB := makeProducer([]int{10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32})

	merged := merge(chA, chB)

	for x := range merged {
		fmt.Printf(" ... merged: %d \n", x)
	}

}

func merge(a, b <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for a != nil || b != nil {
			select {
			case x, ok := <-a:
				// chan is closed ok is set to false
				if !ok {
					fmt.Println("setting chan A to nil")
					// when that is the case, mark the channel as nil so that
					// <-nil is a no-op
					a = nil
					continue
				}
				out <- x
			case x, ok := <-b:
				if !ok {
					fmt.Println("setting chan B to nil")
					b = nil
					continue
				}

				out <- x
			}
		}
	}()
	return out
}

func makeProducer(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for _, x := range nums {
			out <- x
			time.Sleep(time.Duration(rand.Intn(700)) * time.Millisecond)
		}
	}()
	return out
}
