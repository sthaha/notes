package main

import (
	"context"
	"log"
	"net/http"
	"sync"
	"time"
)

func startServer(ctx context.Context, wg *sync.WaitGroup, addr string, handle http.Handler) *http.Server {
	srv := &http.Server{
		Addr:           addr,
		Handler:        handle,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		select {
		case <-ctx.Done():
			log.Printf("ctx done: shutting down server at %q ", srv.Addr)
			// We received an interrupt signal, shut down.
			if err := srv.Shutdown(context.Background()); err != nil {
				// Error from closing listeners, or context timeout:
				log.Printf("Error: server shutdown: %v", err)
			}
		}
	}()

	wg.Add(1)
	defer wg.Done()
	log.Printf("Going to start server at %q ", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
	return srv
}
