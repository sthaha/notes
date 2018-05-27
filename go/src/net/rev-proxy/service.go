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
	log.Printf(">> rootHandler: Got: %s\n", r.URL)
	response := fmt.Sprintf("Service: path: %q\n", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, response)
	io.WriteString(w, response)

	time.Sleep(config.delay)
	io.WriteString(w, "End of response")

	log.Printf("Wrote response for: %s\n", r.URL)
}

func (config serverConfig) ok(w http.ResponseWriter, r *http.Request) {
	log.Printf(">> OK Got: %s\n", r.URL)
	response := fmt.Sprintf(">> server: path: %q\n", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, response)
	log.Printf("Wrote response for: %s\n", r.URL)
}

func (config serverConfig) notFound(w http.ResponseWriter, r *http.Request) {
	log.Printf(">> NOT Found: Got: %s\n", r.URL)
	response := fmt.Sprintf(">> server: Not found: %q\n", r.URL)
	w.WriteHeader(404)
	io.WriteString(w, response)
	log.Printf("Wrote response for: %s\n", r.URL)
}

func startService(
	ctx context.Context, wg *sync.WaitGroup,
	addr string, delay time.Duration) {
	mux := http.NewServeMux()

	config := &serverConfig{delay: delay}
	mux.HandleFunc("/api/ok/", config.ok)
	mux.HandleFunc("/api/notfound/", config.notFound)
	mux.HandleFunc("/api/", config.rootHandler)

	startServer(ctx, wg, addr, mux)
}
