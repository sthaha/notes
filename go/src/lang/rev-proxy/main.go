package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type service struct {
}

func newService() *service {
	return &service{}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Url: %q", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, response)
}

func startService(wg *sync.WaitGroup, ctx context.Context, addr string) {
	mux := newService()
	s := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		select {
		case <-ctx.Done():
			log.Printf("interrupted: shutting down server at %q ", s.Addr)
			// We received an interrupt signal, shut down.
			if err := s.Shutdown(context.Background()); err != nil {
				// Error from closing listeners, or context timeout:
				log.Printf("HTTP server Shutdown: %v", err)
			}
		}
	}()

	wg.Add(1)
	defer wg.Done()
	log.Printf("Going to start server at %q ", s.Addr)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
}

func interruptableContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Printf("--------------------------------------------------")
		log.Printf("interrupted: cancelling context")
		cancel()
	}()
	return ctx
}

func main() {
	var wg sync.WaitGroup
	ctx := interruptableContext()
	go startService(&wg, ctx, ":8080")
	go startService(&wg, ctx, ":8081")
	go startService(&wg, ctx, ":8081")

	<-ctx.Done()
	log.Printf("waiting for servers to terminate")
	wg.Wait()
	log.Printf("All Good .. bye!")
}
