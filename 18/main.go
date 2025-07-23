package main

import (
	"fmt"
	"sync"
)

/*
Вообще не совсем понимаю, что будет являться best practice для использования embed struct или определение через
переменную.
*/

type ConcurrentCounter struct {
	sync.RWMutex
	value int
}

func NewConcurrentCounter() *ConcurrentCounter {
	return &ConcurrentCounter{
		value: 0,
	}
}

func (c *ConcurrentCounter) Increment(wg *sync.WaitGroup) {
	c.Lock()
	defer c.Unlock()
	defer wg.Done()
	c.value++
}

func (c *ConcurrentCounter) Get() int {
	c.RLock()
	defer c.RUnlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	concurrent := NewConcurrentCounter()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go concurrent.Increment(&wg)
	}
	wg.Wait()

	fmt.Println(concurrent.Get())
}
