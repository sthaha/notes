package main

import (
	"fmt"
	"time"
)

type Buffer struct {
	id int
}

func (b *Buffer) process() {
	fmt.Printf("\t... %v: called process \n", b)
	time.Sleep(time.Millisecond * 130)
}

func (b *Buffer) get() {
	fmt.Printf("\t... %v: called get \n", b)
	time.Sleep(time.Millisecond * 220)
}

var available = make(chan Buffer, 10)
var forProcessing = make(chan Buffer)

var currentBuffer int = 1

func worker() {

	for {
		var b Buffer

		select {
		case b = <-available:
			fmt.Printf(" >>> worker: got available buffer: %v\n", b.id)
		default:
			b = Buffer{currentBuffer}
			fmt.Printf(" >>> worker:           new Buffer: %v\n", b.id)
			currentBuffer += 1
		}
		fmt.Printf(" >>> worker: calling get on  buffer: %v\n", b.id)
		b.get()
		forProcessing <- b
	}
}

func controller() {
	for {
		b := <-forProcessing
		fmt.Printf(" <<< controller: got buffer %v \n", b)
		b.process()

		select {
		case available <- b:
			fmt.Printf(" <<< controller: Wrote to available %v \n", b)
		default:
			fmt.Printf(" <<< controller: available buffer full\n")
		}

	}
}

func main() {
	go worker()
	go worker()
	go worker()
	go worker()

	go controller()
	<-time.After(8 * time.Second)
	fmt.Println("Exit ...")
}
