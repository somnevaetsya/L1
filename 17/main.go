package main

import "fmt"

func binarySearch(arr []int, left, right, key int) int {
	mid := 0
	for {
		mid = (left + right) / 2
		if key < arr[mid] {
			right = mid - 1
		} else if key > arr[mid] {
			left = mid + 1
		} else {
			return mid
		}

		if left > right {
			return -1
		}
	}
}

func main() {
	arr := []int{0, 1, 2, 4, 5, 12, 32, 99, 123, 1234, 2345}
	fmt.Println(binarySearch(arr, 0, len(arr)-1, 55))
}
