package main

import "fmt"

type MaxHeap struct {
    items   []int
    maxSize int
}

func NewMaxHeap(maxSize int) *MaxHeap {
    return &MaxHeap{
        make([]int, 0),
        maxSize,
    }
}

func (h *MaxHeap) Push(item int) error {
    if len(h.items) >= h.maxSize {
        return fmt.Errorf("error: %s", "heap is full")
    }

    h.items = append(h.items, item)
    h.upHeapify(len(h.items) - 1)
    return nil
}

func (h *MaxHeap) upHeapify(index int) {
    if index == 0 {
        return
    }

    for h.items[index] > h.items[h.parent(index)] {
        h.swap(index, h.parent(index))
        index = h.parent(index)
    }
}

func (h *MaxHeap) Pop() int {
    max := h.items[0]
    h.items[0] = h.items[len(h.items)-1]
    h.items = h.items[:len(h.items)-1]
    h.downHeapify(0)
    return max
}

func (h *MaxHeap) downHeapify(current int) {
    if current <= h.maxSize && current >= len(h.items)/2 {
        return
    }

    maxIndex := current
    leftChildIndex := current*2 + 1
    rightChildIndex := current*2 + 2

    if leftChildIndex < len(h.items) && h.items[leftChildIndex] > h.items[maxIndex] {
        maxIndex = leftChildIndex
    }

    if rightChildIndex < len(h.items) && h.items[rightChildIndex] > h.items[maxIndex] {
        maxIndex = rightChildIndex
    }

    if current != maxIndex {
        h.swap(current, maxIndex)
        h.downHeapify(maxIndex)
    }
    return
}

func (h *MaxHeap) parent(index int) int {
    return (index - 1) / 2
}

func (h *MaxHeap) swap(i, j int) {
    h.items[i], h.items[j] = h.items[j], h.items[i]
}

func main() {
    arr := []int{3, 4, 2, 6, 5, 8}
    maxHeap := NewMaxHeap(len(arr))
    for _, k := range arr {
        err := maxHeap.Push(k)
        if err != nil {
            fmt.Printf("error: %s", err)
        }
    }

    fmt.Println(maxHeap.items)
}
