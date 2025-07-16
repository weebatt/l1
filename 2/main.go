package main

import (
	"fmt"
	"sync"
)

func main() {
	array := [5]int{2, 4, 6, 8, 10}
	array = CompetitiveSquaring(&array)
	for i, value := range array {
		fmt.Printf("i_%d: %d\n", i, value)
	}
}

func CompetitiveSquaring(array *[5]int) [5]int {
	var wg sync.WaitGroup
	wg.Add(len(array))

	for i := 0; i < len(array); i++ {
		go func() {
			array[i] *= array[i]
			wg.Done()
		}()
	}

	wg.Wait()

	return *array
}
