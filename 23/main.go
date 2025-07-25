package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	removeElement(slice, 3)
}

func removeElement(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])

	fmt.Printf("remove element by index: %d. Current slice: %v\n", index, slice)

	slice[len(slice)-1] = 0

	fmt.Printf("Set last element to 0. Current slice: %v\n", slice)

	slice = slice[:len(slice)-1]

	fmt.Printf("Final slice: %v\n", slice)

	return slice
}
