package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx := withInterruptibleContext(context.Background())

	go startService(ctx, &wg, ":8081", 8*time.Second)
	go startReverseProxy(ctx, &wg, ":8000", "http://localhost:8081")

	<-ctx.Done()
	log.Printf("waiting for servers to terminate")
	wg.Wait()
	log.Printf("All Good .. bye!")
}
