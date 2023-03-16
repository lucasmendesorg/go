/*
	Sort in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>
*/

package main

import "fmt"
import "math/rand"

func swapArrayIndexes(arr []int, indexA int, indexB int) {
		x := arr[indexA]
		y := arr[indexB]
		arr[indexA] = y
		arr[indexB] = x
}

func sortArrayAscending2(arr []int) {
	for done, l := false, len(arr) - 1; done == false; {
		done = true
		for i := 0; i < l; i++ {
			if arr[i + 1] < arr[i] {
				swapArrayIndexes(arr, i, i + 1)
				done = false
			}
		}
	}
}

func sortArrayAscendingRecursive(arr []int, current int, limit int) bool {
	if current == limit {
		return true
	}
	if arr[current + 1] < arr[current] {
		swapArrayIndexes(arr, current, current + 1)
		return sortArrayAscendingRecursive(arr, 0, limit)
	}
	current++
	return sortArrayAscendingRecursive(arr, current, limit)
}

func sortArrayAscending(arr []int) {
	for done, l := false, len(arr) - 2; done == false; {
		done = true
		for i := l; i >= 0; i-- {
			if arr[i + 1] < arr[i] {
				swapArrayIndexes(arr, i, i + 1)
				done = false
			}
		}
	}
}

func sortArrayDescending(arr []int) {
	for done, l := false, len(arr) - 1; done == false; {
		done = true
		for i := 0; i < l; i++ {
			if arr[i + 1] > arr[i] {
				swapArrayIndexes(arr, i, i + 1)
				done = false
			}
		}
	}
}

func createRandomArray(length uint) []int {
	arr := make([]int, length)
	for i := uint(0); i < length; i++ {
		arr[i] = rand.Intn(100)
	}
	return arr
}

func doSortDescending() {
	fmt.Println("--- sortArrayDescending ---")
	arr := createRandomArray(32)
	fmt.Println("before: arr =", arr)
	sortArrayDescending(arr)
	fmt.Println(" after: arr =", arr)
}

func doSortAscending2() {
	fmt.Println("--- sortArrayAscending2 ---")
	arr := createRandomArray(32)
	fmt.Println("before: arr =", arr)
	sortArrayAscending2(arr)
	fmt.Println(" after: arr =", arr)
}

func doSortAscending() {
	fmt.Println("--- sortArrayAscending ---")
	arr := createRandomArray(32)
	fmt.Println("before: arr =", arr)
	sortArrayAscending(arr)
	fmt.Println(" after: arr =", arr)
}

func doSortAscendingRecursive() {
	fmt.Println("--- sortArrayAscendingRecursive ---")
	arr := createRandomArray(32)
	fmt.Println("before: arr =", arr)
	sortArrayAscendingRecursive(arr, 0, len(arr) - 1)
	fmt.Println(" after: arr =", arr)
	
}

func main() {
	doSortAscendingRecursive()
	doSortAscending()
	doSortAscending2()
	doSortDescending()
}
