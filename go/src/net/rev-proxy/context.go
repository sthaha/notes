package main

import (
	"context"
	"log"
	"os"
	"os/signal"
)

func cancelOnSigInt(cancel context.CancelFunc) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
	log.Printf("--------------------------------------------------")
	log.Printf("interrupted: cancelling context")
	cancel()
}

func withInterruptibleContext(parent context.Context) context.Context {
	ctx, cancel := context.WithCancel(parent)
	go cancelOnSigInt(cancel)
	return ctx
}
