/*
	Heap in Go
		Written by Lucas Mendes <lucas AT lucasmendes.org>

	Time complecity = O(h) => O(log n), for h => Height of the heap
*/

package main

import (
	"errors"
	"fmt"
)

type MaxHeap struct {
	array []int
}

func (h *MaxHeap) swap(indexA, indexB int) {
	h.array[indexA], h.array[indexB] = h.array[indexB], h.array[indexA]
}

func (h *MaxHeap) right(index int) int {
	return index*2 + 2
}

func (h *MaxHeap) left(index int) int {
	return index*2 + 1
}

func (h *MaxHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) heapifyDown(index int) {
	lastIndex := len(h.array) - 1
	left, right := h.left(index), h.right(index)
	childToCompare := 0
	for left <= lastIndex {
		if left == lastIndex {
			childToCompare = left
		} else if h.array[left] > h.array[right] {
			childToCompare = left
		} else {
			childToCompare = right
		}
		if h.array[index] > h.array[childToCompare] {
			break
		}
		h.swap(index, childToCompare)
		index = childToCompare
		left, right = h.left(index), h.right(index)
	}
}

func (h *MaxHeap) heapifyUp(index int) {
	for parentIndex := h.parent(index); h.array[parentIndex] < h.array[index]; parentIndex = h.parent(index) {
		h.swap(parentIndex, index)
		index = parentIndex
	}
}

func (h *MaxHeap) Extract() (int, error) {
	arrayLength := len(h.array)
	if arrayLength == 0 {
		return 0, errors.New("cannot extract key from empty array")
	}
	extracted := h.array[0]
	lastIndex := arrayLength - 1
	h.array[0] = h.array[lastIndex]
	h.array = h.array[:lastIndex]
	h.heapifyDown(0)
	return extracted, nil
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.heapifyUp(len(h.array) - 1)
}

func main() {
	m := MaxHeap{}
	fmt.Println("m =", m)
	for _, v := range []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100} {
		m.Insert(v)
	}
	fmt.Println("m =", m)
	for i := 0; i < 3; i++ {
		fmt.Printf("--- Extracting test %d ---\n", i)
		extracted, err := m.Extract()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("extracted =", extracted)
		fmt.Println("m =", m)
	}
}
