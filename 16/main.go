package main

func main() {
	arr := []int{2, 5, 8, 1, 3, 4}
	arr = quickSort(arr)
}

func quickSort(arr []int) []int {
	return quickSortHelper(arr, 0, len(arr)-1)
}

func quickSortHelper(arr []int, first, last int) []int {
	if first < last {
		pivotIndex := partition(arr, first, last)
		quickSortHelper(arr, first, pivotIndex-1)
		quickSortHelper(arr, pivotIndex+1, last)
	}

	return arr
}

func partition(arr []int, first, last int) int {
	pivot := arr[first]
	left, right := first+1, last

	for done := false; done != true; {
		for left <= right && arr[left] <= pivot {
			left++
		}

		for arr[right] >= pivot && right >= left {
			right--
		}

		if right < left {
			done = true
		} else {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[first], arr[right] = arr[right], arr[first]
	return right
}
