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

var delay = 8

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("<<< Service: Got: %s\n", r.URL)
	response := fmt.Sprintf("Server: Url: %q\n", r.URL)

	log.Printf("<<< Sleeping for %d seconds\n", delay)
	time.Sleep(time.Duration(delay) * time.Second)
	delay--
	w.WriteHeader(200)
	io.WriteString(w, response)
	log.Printf("<<< Wrote response for: %s\n", r.URL)
}

func startService(ctx context.Context, wg *sync.WaitGroup, addr string, basePath string) {
	mux := http.NewServeMux()
	mux.HandleFunc(basePath, rootHandler)

	startServer(ctx, wg, addr, mux)
}
