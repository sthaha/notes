package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
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

func startService(ctx context.Context, addr string) {
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

	log.Printf("Going to start server at %q ", s.Addr)

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
}

func cancelContextOnInterrupt(cancel context.CancelFunc) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Printf("--------------------------------------------------")
	log.Printf("interrupted: cancelling context")
	cancel()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go cancelContextOnInterrupt(cancel)
	go startService(ctx, ":8080")
	go startService(ctx, ":8081")
	go startService(ctx, ":8082")

	<-ctx.Done()
	log.Printf(".... waiting 2 seconds for servers to terminate")
	time.Sleep(2 * time.Second)
}
