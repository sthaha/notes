package main

import (
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type workerPool struct {
	workers   chan bool
	stop      chan bool
	jobsWrite chan<- any
	jobsRead  <-chan any
}

func newWorkerPool(n int) *workerPool {
	w, r := unbufferedChan()
	return &workerPool{
		workers:   make(chan bool, n),
		stop:      make(chan bool, n),
		jobsWrite: w,
		jobsRead:  r,
	}
}

func (wp *workerPool) Shutdown() {
	close(wp.stop)
}

func (wp *workerPool) Run() {
	var wg sync.WaitGroup
	go func() {
		index := 0
		for {

			wp.workers <- true
			wg.Add(1)
			index = (index + 1) % cap(wp.workers)

			go func(num int) {
				log.Printf("[%d] started worker ...", num)
				defer wg.Done()
				select {
				case <-wp.stop:
					log.Printf("[%d] stopping workers ...", num)
					return
				case job := <-wp.jobsRead:
					fn := job.(func())
					fn()
				}

				<-wp.workers
			}(index)
		}

	}()

	<-wp.stop

	wg.Wait()
}

func (wp *workerPool) Queue(j func()) {
	wp.jobsWrite <- j
}

func main() {
	wp := newWorkerPool(5)
	stopOnInterrupt(wp)
	//pushJobs(wp)
	wp.Run()
}

func pushJobs(wp *workerPool) {
	go func() {
		for i := 0; i < 100; i++ {

			d := rand.Intn(2000)
			time.Sleep(time.Duration(d) * time.Millisecond)
			wp.Queue(func() {
				wait := rand.Intn(10)
				log.Printf(".... going to run for %d", wait)
				time.Sleep(time.Duration(wait) * time.Second)
				log.Printf(".... done running for %d", wait)
			})

		}

	}()
}

func stopOnInterrupt(wp *workerPool) {
	osSigs := make(chan os.Signal, 1)
	signal.Notify(osSigs, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("Waiting for signal ...")
	go func() {
		sig := <-osSigs
		log.Printf("Recieved signal %v ; stopping workers", sig)
		wp.Shutdown()
	}()

}
