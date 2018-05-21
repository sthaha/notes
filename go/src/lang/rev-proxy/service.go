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

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service: Got: %s\n", r.URL)
	response := fmt.Sprintf("Server: Url: %q\n", r.URL)
	time.Sleep(5 * time.Second)
	w.WriteHeader(200)
	io.WriteString(w, response)
}

func startService(ctx context.Context, wg *sync.WaitGroup, addr string, basePath string) {
	mux := http.NewServeMux()
	mux.HandleFunc(basePath, rootHandler)

	startServer(ctx, wg, addr, mux)
}
