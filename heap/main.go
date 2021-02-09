package main

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

func (h *MaxHeap) maxHeapifyUp(index int) {

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
	
}