package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ждем две секунды")
	start := time.Now()
	Sleep(3)
	duration := time.Since(start)
	fmt.Printf("Подождали %.5f секунды\n", duration.Seconds())
}

func Sleep(duration int) {
	<-time.After(time.Duration(duration) * time.Second)
}
