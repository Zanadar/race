package main

import (
	"fmt"
	"sync"
	"time"
)

type intCounter struct {
	counter int64
	mu      sync.Mutex
}

func (c *intCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
	return
}

func (c *intCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.counter
}

func main() {
	counter := intCounter{}

	for i := 0; i < 100; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				counter.Inc()
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(counter.Value())

}
