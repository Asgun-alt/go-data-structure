package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.MaxHeapifyUp(len(h.array)-1)
}

// Extract return the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int{
	extracted := h.array[0]
	last := len(h.array)-1

	// check if array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract 0 array lenght")
		return -1
	}

	// take out the last index and put it in the root
	h.array[0] = h.array[last] // get the last index and put it in index 0
	h.array = h.array[:last]

	h.MaxHeapifyDown(0)
	
	return extracted
}

// Swap if the index is larger than the parent
// MaxHapifyUp will heapify from the bottom top
func (h *MaxHeap) MaxHeapifyUp(index int){
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}
// MaxHeapifyDown will heapify top to bottom
func (h *MaxHeap) MaxHeapifyDown(index int){
	lastIndex := len(h.array)-1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex{
		if l == lastIndex{ // when left child is the  only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare]{
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// Get the parent index
func parent(i int) int{
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int{
	return 2 * i + 1
}

// Get the right child index
func right(i int) int{
	return 2 * i + 2
}

// swap keys in the array
func (h * MaxHeap) swap(idx1, idx2 int){
	h.array[idx1], h.array[idx2] = h.array[idx2], h.array[idx1]

}

func main() {
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 40 ,50, 70, 12, 6, 2}
	for _, val := range buildHeap{
		m.Insert(val)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}