package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomDuration() time.Duration {
	return time.Duration(rand.Intn(1e8))
}

func retry(f func()) <-chan bool {
	c := make(chan bool)

	go
}

func main() {
	start := time.Now()

	var t *time.Timer

	t = time.AfterFunc(randomDuration(), func() {
		next := randomDuration()
		fmt.Println(time.Now().Sub(start), "next: ", next)
		if !t.Stop() {
			<-t.C
		}
		t.Reset(next) // READ doc to understand why this is data race
	})
	time.Sleep(5 * time.Second)
}
