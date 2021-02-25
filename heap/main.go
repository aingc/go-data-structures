package main

import (
	"fmt"
)

// a heap is a data structure that can be expressed as a complete tree
// a complete tree means that all the levels in the tree are full except for the lowest
// level where some nodes can be empty but have all its nodes to the left

// strictly speaking heaps are actually an array that operates like a tree
// with a parent index you can get both left child index and right child index
// via: (parentIdx * 2) + 1 for left child index, and (parentIdx * 2) + 2 for right child index

// When inserting keys in a heap we add to the bottom-right most place
// After placing we need to heapify making sure that the parent key is larger than its
// children, compare children to parent and swap accordingly until they get to its correct
// place, also needed to do after extracting a key

// When extracting the root which is the max value, we replace the root with the
// last node, then heapify downwards from the root and decide which of the 3 numbers
// which is the greatest value and that should be the parent, and then keep swapping
// downwards accordingly

// Heapify affects the time complexity the most O(h) where h is the height of the tree
// Replace h with logn. For Heap insert and extract is O(logn)

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array)-1

	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	// set last index value to root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	// When parent is smaller than current index
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array)-1
	l,r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex {
		// Define child to compare
		// When there is only one child or the 
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {  // left child is larger
			childToCompare = l
		} else { // When right child is larger
			childToCompare = r
		}
		// Compare array value of current idx to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l,r = left(index), right(index)
		} else {
			return
		}
	}
}

// Getting parent index is (idx - 1)/2
func parent(i int) int {
	return (i-1)/2
}

// Get left child index: (i*2) + 1
func left(i int) int {
	return i*2 + 1
}

// Get right child index: (i*2) + 2
func right(i int) int {
	return i*2 + 2
}

// Swap keys in slice
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1] // this swapping works in Go
}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}