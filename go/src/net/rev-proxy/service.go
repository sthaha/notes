package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

type serverConfig struct {
	delay time.Duration
}

func (config serverConfig) rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service: Got: %s\n", r.URL)
	response := fmt.Sprintf("Hello from server: path: %q\n", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, response)

	time.Sleep(config.delay)
	io.WriteString(w, "End of response")

	log.Printf("Wrote response for: %s\n", r.URL)
}

func startService(ctx context.Context, wg *sync.WaitGroup, addr string, basePath string, delay time.Duration) {
	mux := http.NewServeMux()
	config := &serverConfig{delay: delay}
	mux.HandleFunc(basePath, config.rootHandler)

	startServer(ctx, wg, addr, mux)
}
