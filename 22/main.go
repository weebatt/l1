package main

import "fmt"

func main() {
	a := int64(2000000)
	b := int64(3000000)

	fmt.Println(add(a, b))
	fmt.Println(sub(a, b))
	fmt.Println(mul(a, b))
	fmt.Println(div(a, b))
}

func add(a, b int64) int64 {
	return a + b
}

func sub(a, b int64) int64 {
	return a - b
}

func mul(a, b int64) int64 {
	return a * b
}

func div(a, b int64) float64 {
	if b != 0 {
		return float64(a) / float64(b)
	}
	return -1
}
