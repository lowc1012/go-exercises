package main

import "fmt"

type MinHeap struct {
	items   []int
	maxSize int
}

func NewMinHeap(maxSize int) *MinHeap {
	return &MinHeap{
		items:   make([]int, 0),
		maxSize: maxSize,
	}
}

func (h *MinHeap) leaf(index int) bool {
	if index < h.maxSize && index >= (len(h.items)/2) {
		return true
	}
	return false
}

func (h *MinHeap) parent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) leftChild(index int) int {
	return index*2 + 1
}

func (h *MinHeap) rightChild(index int) int {
	return index*2 + 2
}

// Push pushes the element x onto the heap.
// The complexity is O(log n)
func (h *MinHeap) Push(item int) error {
	if len(h.items) >= h.maxSize {
		return fmt.Errorf("error: %s", "heap is full")
	}
	h.items = append(h.items, item)
	h.upHeapify(len(h.items) - 1)
	return nil
}

// upHeapify is used when we insert a new element to a heap. When inserting a new
// element, we add it at the bottom of the heap tree, and move up the tree while comparing
// to the current parent element and swapping if needed.
func (h *MinHeap) upHeapify(index int) {
	if index == 0 {
		return
	}

	for h.items[index] < h.items[h.parent(index)] {
		fmt.Printf("index(%d) < parent (%d) -> swap\n", h.items[index], h.items[h.parent(index)])
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *MinHeap) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

// Pop removes and returns the minimum element (according to Less) from the heap.
// The complexity is O(log n)
func (h *MinHeap) Pop() int {
	min := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.downHeapify(0)
	return min
}

// downHeapify is used when we remove the top element from a heap
// Removal of an element is done by swapping the top element with
// the last element at the bottom of the tree, removing the last element,
// and then heapfying the new top element down to maintain the heap property.
func (h *MinHeap) downHeapify(current int) {
	if h.leaf(current) {
		return
	}

	minIndex := current
	leftChildIndex := h.leftChild(current)
	rightChildIndex := h.rightChild(current)

	if leftChildIndex < len(h.items) && h.items[leftChildIndex] < h.items[minIndex] {
		minIndex = leftChildIndex
	}

	if rightChildIndex < len(h.items) && h.items[rightChildIndex] < h.items[minIndex] {
		minIndex = rightChildIndex
	}

	if minIndex != current {
		h.swap(minIndex, current)
		h.downHeapify(minIndex)
	}
	return
}

func (h *MinHeap) buildMinHeap() {
	for i := len(h.items)/2 - 1; i >= 0; i-- {
		h.downHeapify(i)
	}
}

func main() {
	inputArray := []int{7, 3, 4, 2, 6}
	minHeap := NewMinHeap(len(inputArray))
	for i := 0; i < len(inputArray); i++ {
		err := minHeap.Push(inputArray[i])
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println(minHeap.items)

	for i := 0; i < len(inputArray); i++ {
		fmt.Println(minHeap.Pop())
	}
}
