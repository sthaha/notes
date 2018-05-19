package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type service struct {
}

func NewService() *service {
	return &service{}
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf("Url: %q", r.URL)
	w.WriteHeader(200)
	io.WriteString(w, response)
}

func startService() {
	mux := NewService()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

}

func main() {
	startService()

}
