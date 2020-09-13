package main

import (
	"sync"
	"testing"
	"time"
)

func TestUnbuffered(t *testing.T) {
	in, out := unbufferedChan()

	wg := sync.WaitGroup{}
	push(&wg, t, in)
	pop(&wg, t, out)
	wg.Wait()
}

func push(wg *sync.WaitGroup, t *testing.T, in chan<- any) {
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			t.Logf("push %d", i)
			in <- i
			time.Sleep(20 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
		for i := 100; i < 110; i++ {
			t.Logf("push %d", i)
			in <- i
			time.Sleep(20 * time.Millisecond)
		}

		close(in)
	}()
}

func pop(wg *sync.WaitGroup, t *testing.T, out <-chan any) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for x := range out {
			t.Logf("popped %d", x)
			time.Sleep(100 * time.Millisecond)
		}
	}()

}
