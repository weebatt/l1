package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	_, _ = fmt.Scan(&n)

	var i int
	_, _ = fmt.Scan(&i)

	if ValidateSpecificBit(i) {
		result := InversingSpecificBit(i, n)
		fmt.Println(result)
	} else {
		fmt.Println("You entered specific bit not in [0, 64) diapason")
	}
}

func InversingSpecificBit(i int, n int64) int64 {
	return n ^ int64(math.Pow(2, float64(i)))
}

func ValidateSpecificBit(i int) bool {
	return i >= 0 && i < 64
}
