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

	go startService(ctx, &wg, ":8081", "/srv1/", 8*time.Second)
	go startService(ctx, &wg, ":8082", "/srv2/", 4*time.Second)
	go startService(ctx, &wg, ":8083", "/srv3/", 3*time.Second)
	go startReverseProxy(ctx, &wg, ":8000", "http://localhost:8081/srv1")

	<-ctx.Done()
	log.Printf("waiting for servers to terminate")
	wg.Wait()
	log.Printf("All Good .. bye!")
}
