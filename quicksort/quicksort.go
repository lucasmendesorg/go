/*
	QuickSort in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "fmt"

func QuickSort(arr []int, begin int, end int) []int {
	var partition = func(arr []int, begin int, end int) int {
		pivot := arr[end-1]
		for i := begin; i < end; i++ {
			if arr[i] <= pivot {
				arr[i], arr[begin] = arr[begin], arr[i]
				begin += 1
			}
		}
		return begin - 1
	}
	if end == -1 {
		end = len(arr)
	}
	if begin < end {
		sp := partition(arr, begin, end)
		QuickSort(arr, begin, sp)
		QuickSort(arr, sp+1, end)
	}
	return arr
}

func main() {
	arr := []int{8, 5, 12, 55, 3, 7, 82, 44, 35, 25, 41, 29, 17}
	fmt.Println("            arr =", arr)
	QuickSort(arr, 0, len(arr))
	fmt.Println("quicksorted arr =", arr)
}
