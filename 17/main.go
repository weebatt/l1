package main

import "fmt"

func main() {
	var target int
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	_, _ = fmt.Scan(&target)

	fmt.Println(binarySearch(arr, target, 0, len(arr)-1))
}

func binarySearch(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}

	mid := left + (right-left)/2

	fmt.Println("target:", target, "left:", left, "right:", right, "mid:", mid)

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return binarySearch(arr, target, mid+1, right)
	}
	return binarySearch(arr, target, left, mid-1)
}
