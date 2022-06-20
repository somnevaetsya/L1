package main

import "fmt"

func partition(arr []int, start, end int) ([]int, int) {
	// выбираем конечную точку
	pivot := arr[end]
	i := start
	for j := start; j < end; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[end] = arr[end], arr[i]
	return arr, i
}

func quickSort(arr []int, start, end int) []int {
	if start < end {
		var pivot int
		arr, pivot = partition(arr, start, end)
		arr = quickSort(arr, start, pivot-1)
		arr = quickSort(arr, pivot+1, end)
	}
	return arr
}

func main() {
	arr := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
	arr = quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
