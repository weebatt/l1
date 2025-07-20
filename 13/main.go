package main

import "fmt"

func main() {
	a := 1
	b := 2
	fmt.Printf("Input a: %d and b:%d\n", a, b)

	// By arithmetics operations
	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("After arithmetics a: %d and b:%d\n", a, b)

	// By XOR
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf("After XOR a: %d and b:%d\n", a, b)
}
