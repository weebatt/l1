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

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	var n int
	_, _ = fmt.Scan(&n)

	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(2)

	// Sender
	go Sender(ch, ctx, &wg)

	// Receiver
	go Receiver(ch, ctx, &wg)

	select {
	case <-time.After(time.Duration(n) * time.Second):
		fmt.Println("Time has expired")
		stop()
		return
	case <-ctx.Done():
	}

	wg.Wait()

	close(ch)
}

func Sender(ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Androids producing factory goroutine exit\n")
			wg.Done()
			return
		case ch <- fmt.Sprintf("Android %d was succesfully produced", rand.Intn(1000)):
			time.Sleep(500 * time.Millisecond)
		}

	}
}

func Receiver(ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Androids consumer goroutine exit\n")
			wg.Done()
			return
		case msg := <-ch:
			fmt.Println(msg + " and consumed!")
		}
	}
}
