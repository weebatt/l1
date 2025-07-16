package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)

	ch := make(chan float64)

	// Generating infinite data to chanel
	go Generator(ch)

	// Creating workers to receive data from infinite chanel
	Receivers(n, ch)

	/*
		Wait without any cases, like inf cycle for main goroutine. Another variant is using WaitGroup.
		To interrupt use hot key cmd + c / ctrl + c to push SIGINT
	*/
	select {}
}

func Generator(ch chan float64) {
	for {
		ch <- rand.Float64() * 10000
	}
}

func Receivers(amount int, ch chan float64) {
	for i := 0; i < amount; i++ {
		go func() {
			for {
				msg := <-ch
				fmt.Printf("i_%d: %f\n", i, msg)
				time.Sleep(500 * time.Millisecond)
			}
		}()
	}
}
