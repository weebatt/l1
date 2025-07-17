package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeMap struct {
	mutex sync.Mutex
	m     map[int]string
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]string),
	}
}

func (s *SafeMap) Set(key int, value string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key int) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.m[key]
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("USING THE SYNC MAP")
	var concMap sync.Map

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			concMap.Store(i, fmt.Sprintf("WBCHECK%d", i))
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Printf("i_%d: ", i)
		fmt.Println(concMap.Load(i))
	}

	fmt.Println("USING THE CUSTOM SAFE MAP")
	safeMap := NewSafeMap()

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			safeMap.Set(i, fmt.Sprintf("WBCHECK%d", i))
			time.Sleep(300 * time.Millisecond)
		}()
	}
	wg.Wait()

	for i := 0; i < 10; i++ {
		fmt.Println(fmt.Sprintf("i_%d:", i), safeMap.Get(i))
	}
}
