package main

import (
	"context"
	"log"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ctx := withInterruptibleContext(context.Background())

	go startService(ctx, &wg, ":8081", "/srv1/")
	go startService(ctx, &wg, ":8082", "/srv2/")
	go startService(ctx, &wg, ":8083", "/srv3/")
	go startReverseProxy(ctx, &wg, ":8000", "http://localhost:8081/srv1")

	<-ctx.Done()
	log.Printf("waiting for servers to terminate")
	wg.Wait()
	log.Printf("All Good .. bye!")
}
