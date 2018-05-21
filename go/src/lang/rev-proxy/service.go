package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Service: Got: %s\n", r.URL)
	response := fmt.Sprintf("Server: Url: %q\n", r.URL)
	io.WriteString(w, response)
	w.WriteHeader(200)
}

func startService(ctx context.Context, wg *sync.WaitGroup, addr string, basePath string) {
	mux := http.NewServeMux()
	mux.HandleFunc(basePath, rootHandler)

	startServer(ctx, wg, addr, mux)
}
