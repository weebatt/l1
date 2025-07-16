package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx := context.Background()

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup

	var n int
	_, _ = fmt.Scan(&n)

	ch := make(chan float64)

	// Generating infinite data to chanel
	wg.Add(1)
	go Generator(ctx, &wg, ch)

	// Creating workers to receive data from infinite chanel
	wg.Add(n)
	Receivers(ctx, &wg, n, ch)

	// Wait SIGINT or SIGTERM
	<-ctx.Done()

	// Wait while all goroutines finished
	wg.Wait()
}

func Generator(ctx context.Context, wg *sync.WaitGroup, ch chan float64) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("quit generator")
			wg.Done()
			return
		case ch <- rand.Float64() * 10000:
		}
	}
}

func Receivers(ctx context.Context, wg *sync.WaitGroup, amount int, ch chan float64) {
	for i := 0; i < amount; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					fmt.Printf("quit receiver number %d\n", i)
					wg.Done()
					return
				case msg := <-ch:
					fmt.Printf("i_%d: %f\n", i, msg)
					time.Sleep(500 * time.Millisecond)
				}
			}
		}()
	}
}
