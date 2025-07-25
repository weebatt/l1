package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ждем две секунды")
	Sleep(2)
	fmt.Println("Подождали две секунды")
}

func Sleep(duration int) {
	<-time.After(time.Duration(duration) * time.Second)
}
