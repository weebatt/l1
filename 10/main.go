package main

import (
	"fmt"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[int][]float64
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int][]float64),
	}
}

func (s *SafeMap) Set(key int, value float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = append(s.m[key], value)
}

func main() {
	var wg sync.WaitGroup

	list := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	smap := NewSafeMap()

	wg.Add(1)
	go Forecaster(list, smap, &wg)
	wg.Wait()

	for k, v := range smap.m {
		fmt.Printf("%d: %.1f\n", k, v)
	}
}

func Forecaster(list []float64, smap *SafeMap, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range list {
		smap.Set(int(v)/10*10, v)
	}
}
