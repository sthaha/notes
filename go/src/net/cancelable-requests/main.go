package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

func withTimeoutContext(parent context.Context, duration time.Duration) context.Context {
	ctx, cancel := context.WithCancel(parent)
	time.AfterFunc(duration, func() {
		log.Printf(" ... calling cancel %v", cancel)
		cancel()
	})
	return ctx
}

func clientTimeout(ctx context.Context) {

	retry := true

	go func() {
		select {
		case <-ctx.Done():
			log.Printf("Stopping client retry")
			retry = false
		}
	}()

	for retry {
		req, _ := http.NewRequest("GET", "http://localhost:8080/srv1/foo/bar", nil)
		req = req.WithContext(withTimeoutContext(context.Background(), 3*time.Second))
		log.Printf(" ... context: %v\n", req.Context())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf(">>> Errrrrr :%v", err)
			continue
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf(">>> Body: %s", body)
		return
	}
}

func clientTimeoutInline(ctx context.Context, d time.Duration) {

	for {
		select {
		case <-ctx.Done():
			log.Printf("Stopping client retry")
			return
		default:
		}

		req, _ := http.NewRequest("GET", "http://localhost:8080/srv1/foo/bar", nil)
		req = req.WithContext(
			withTimeoutContext(req.Context(), d),
		)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf(">>> Errrrrr :%v", err)
			continue
		}

		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		log.Printf(">>> Body: %s", body)
		return
	}
}

func main() {
	var wg sync.WaitGroup
	ctx := withInterruptibleContext(context.Background())

	go startService(ctx, &wg, ":8080", "/srv1/")
	go clientTimeoutInline(ctx, 1*time.Second)
	go clientTimeoutInline(ctx, 2*time.Second)
	go clientTimeoutInline(ctx, 3*time.Second)

	<-ctx.Done()
	time.Sleep(5 * time.Second)
	log.Printf("waiting for servers to terminate")
	wg.Wait()
	log.Printf("All Good .. bye!")
}
