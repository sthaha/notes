package main

import (
	"fmt"
)

type MyCustomError struct {
	code int
}

func (e *MyCustomError) Error() string {
	return fmt.Sprintf("Bad code %d", e.code)
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	myErr := &MyCustomError{128}
	fmt.Println("My custom error:", myErr)

	if _, err2 := f2(42); err2 != nil {
		fmt.Println("My custom error:", err2)
	}
}
