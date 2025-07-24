package main

import "fmt"

func main() {
	arr2reverse := []rune("åß∂ƒ©˙hello")

	for i := 0; i < len(arr2reverse)/2; i++ {
		j := len(arr2reverse) - 1 - i
		arr2reverse[i], arr2reverse[j] = arr2reverse[j], arr2reverse[i]
	}

	fmt.Println(string(arr2reverse))
}
