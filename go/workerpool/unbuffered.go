package main

import (
	"log"
	"sync"
)

type any interface{}

type ch struct {
	in      chan any // user writes and closes this
	out     chan any // user can read from this
	buffer  []any
	mutex   sync.Mutex // protect buffer
	present chan bool  // signals if buffer is present
	closed  chan bool  // signals if user closes the in channel
}

func unbufferedChan() (chan<- any, <-chan any) {
	c := ch{
		in:      make(chan any),
		out:     make(chan any),
		present: make(chan bool),
		closed:  make(chan bool),
		buffer:  []any{},
	}

	go c.run()
	return c.in, c.out
}

func (c *ch) run() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer func() {
			log.Printf(".... read closed")
			wg.Done()
		}()
		for {
			log.Printf("... ... reading channels: In")

			select {
			case v, ok := <-c.in:
				if !ok {
					log.Printf("... ... use closed exit read")
					c.in = nil
					close(c.present)
					close(c.closed)
					return
				}

				c.push(v)
			}
		}
	}()

	go func() {
		defer wg.Done()
		defer func() {
			log.Printf(".... closing write")
			close(c.out)
		}()

		for {
			v, length := c.pop()
			log.Printf("....... popped %v : len: %d", v, length)

			if length == 0 {
				log.Printf("....... waiting for data or close")
				select {
				case <-c.present:
					continue
				case <-c.closed:
					if c.length() == 0 {
						return
					}
					continue
				}

			}
			c.out <- v
		}
	}()
	wg.Wait()
	log.Printf("........ outside run")
}

func (c *ch) length() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return len(c.buffer)
}

func (c *ch) clear() {
	c.mutex.Lock()
	c.buffer = []any{}
	c.mutex.Unlock()
}

func (c *ch) push(x any) {
	log.Printf("....... pushing %v", x)

	c.mutex.Lock()
	defer func() {
		log.Printf("....... unlocking mutex")
		c.mutex.Unlock()
	}()
	c.buffer = append(c.buffer, x)

	select {
	case c.present <- true:
		log.Printf("....... unblocking - pushed")
	default:
		log.Printf("....... unblocked already present")
	}
}

func (c *ch) pop() (any, int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	length := len(c.buffer)
	if length == 0 {
		return nil, length
	}

	var first any
	first, c.buffer = c.buffer[0], c.buffer[1:]
	return first, length
}
