package main

import (
	"context"
	"fmt"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

/*
Используя time.Sleep() в main() будет продемонстрировано завершение работы горутин разными подходами.
*/

var shutdown atomic.Bool

func main() {
	var wg sync.WaitGroup
	done := make(chan bool)

	fmt.Println("Started amount of goroutines is ", runtime.NumGoroutine())

	// Stops by closing chanel
	ch1 := make(chan struct{})

	wg.Add(1)

	go PrinterOne(ch1, &wg)
	fmt.Println("PrinterOne started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	close(ch1)
	wg.Wait()

	fmt.Println("PrinterOne exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by termination signal
	ch2 := make(chan bool)

	wg.Add(1)

	go PrinterTwo(ch2, &wg)
	fmt.Println("PrinterTwo started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	ch2 <- true
	wg.Wait()

	fmt.Println("PrinterTwo exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by shared variables
	stopVar := 1

	wg.Add(1)

	go PrinterThree(&stopVar, &wg)
	fmt.Println("PrinterThree started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	stopVar = 0
	wg.Wait()

	fmt.Println("PrinterThree exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by canceling context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)

	go PrinterFour(ctx, &wg)
	fmt.Println("PrinterFour started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()

	fmt.Println("PrinterFour exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by runtime.Goexit()
	wg.Add(1)
	go PrinterFive(&wg)
	fmt.Println("PrinterFive started... Current goroutines number is ", runtime.NumGoroutine())

	wg.Wait()
	fmt.Println("PrinterFive exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by atomic flag
	shutdown.Store(false)

	wg.Add(1)
	go PrinterSix(&wg)
	fmt.Println("PrinterSix started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	shutdown.Store(true)
	wg.Wait()
	fmt.Println("PrinterSix exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	// Stops by catching system signals
	ctx, cancel = signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	wg.Add(1)
	go PrinterSeven(ctx, &wg, done)
	fmt.Println("PrinterSeven started... Current goroutines number is ", runtime.NumGoroutine())
	time.Sleep(3 * time.Second)

	cancel()
	wg.Wait()
	fmt.Println("PrinterSeven exit successfully! Current goroutines number is ", runtime.NumGoroutine())

	/*
		Как я понял еще 2 горутины создаются для отслеживания сигналов системы и их можно закрыть, если создавать
		отдельный канал под получения сигналов от системы.
	*/

	<-done
	fmt.Println("Final amount of goroutines is ", runtime.NumGoroutine())
}

func PrinterOne(ch chan struct{}, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			wg.Done()
			return
		default:
			fmt.Println("Hello from PrinterOne")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func PrinterTwo(ch chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case <-ch:
			wg.Done()
			return
		default:
			fmt.Println("Hello from PrinterTwo")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func PrinterThree(stopVar *int, wg *sync.WaitGroup) {
	for {
		if *stopVar == 0 {
			wg.Done()
			return
		}
		fmt.Println("Hello from PrinterThree")
		time.Sleep(500 * time.Millisecond)
	}
}

func PrinterFour(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			return
		default:
			fmt.Println("Hello from PrinterFour")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func PrinterFive(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Println("Hello from PrinterFive")
		time.Sleep(500 * time.Millisecond)
	}

	runtime.Goexit()
}

func PrinterSix(wg *sync.WaitGroup) {
	for {
		if shutdown.Load() {
			wg.Done()
			return
		}
		fmt.Println("Hello from PrinterSix")
		time.Sleep(500 * time.Millisecond)
	}
}

func PrinterSeven(ctx context.Context, wg *sync.WaitGroup, done chan<- bool) {
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			done <- true
			return
		default:
			fmt.Println("Hello from PrinterSeven")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
