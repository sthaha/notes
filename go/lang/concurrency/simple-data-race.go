package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var (
	count = 0
	wg    sync.WaitGroup
	m     = sync.Mutex{}
)

func incrementCount(ctx context.Context, who string) {
	defer wg.Done()

	inc := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println(who, inc, " ... exiting")
			return
		case <-time.After(400 * time.Millisecond):
			inc += 1
			m.Lock()
			fmt.Println(who, inc, ":   now  :", count)
			count += 1
			fmt.Println(who, inc, ":   inc  :", count)
			m.Unlock()
		}
	}
}

func cancel() {
	fmt.Println(" .... cancel")
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(3)
	go incrementCount(ctx, "A")
	go incrementCount(ctx, "                     B")
	go incrementCount(ctx, "                                           C")

	time.AfterFunc(3*time.Second, cancel)
	wg.Wait()

	fmt.Println("  ... end: ", count)
}
