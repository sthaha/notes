package main

import (
	"log"
	"sync"
	"time"
)

var calls = 0

func whatToRun() {
	calls++
	log.Printf("whatToRun: call: %d \n", calls)
	time.Sleep(3 * time.Second)
	log.Printf("whatToRun: call: %d  - DONE\n", calls)
	calls--
}

var once sync.Once

func onlyOneCall(caller int) {

	log.Printf("Current: %d \n", caller)
	once.Do(whatToRun)
	log.Printf("Current: %d - DONE\n", caller)
}

func main() {
	for i := 1; i <= 10; i++ {
		go onlyOneCall(i)
		time.Sleep(200 * time.Millisecond)
	}
	time.Sleep(5 * time.Second)
}
