package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	wg.Add(2)

	out := Generator(array, &wg)
	squared := Handler(out, &wg)

	wg.Wait()

	for xx := range squared {
		fmt.Println(xx)
	}
}

func Generator(array []int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	out := make(chan int)

	go func() {
		for i := 0; i < len(array); i++ {
			out <- array[i]
			fmt.Printf("index: %d, value: %d\n", i, array[i])
			time.Sleep(500 * time.Millisecond)
		}
		close(out)
	}()

	return out
}

func Handler(in chan int, wg *sync.WaitGroup) chan int {
	defer wg.Done()
	out := make(chan int)

	go func() {
		for x := range in {
			out <- x * x
		}
		close(out)
	}()

	return out
}
