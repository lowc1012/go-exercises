// Example (LeetCode 215) Kth Largest Element in an Array

package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	return quickSelect(&nums, len(nums)-k, 0, len(nums)-1)
}

func partition(nums *[]int, start, end, pivotIndex int) int {
	pivotValue := (*nums)[pivotIndex]

	// Move pivot to the end
	(*nums)[end], (*nums)[pivotIndex] = (*nums)[pivotIndex], (*nums)[end]

	storeIndex := start

	// Move smaller values to left of the pivot
	for i := start; i < end; i++ {
		if (*nums)[i] < pivotValue {
			(*nums)[storeIndex], (*nums)[i] = (*nums)[i], (*nums)[storeIndex]
			storeIndex++
		}
	}

	// Move pivot to its correct position
	(*nums)[storeIndex], (*nums)[end] = (*nums)[end], (*nums)[storeIndex]
	return storeIndex
}

func quickSelect(nums *[]int, k, start, end int) int {

	if start == end {
		return (*nums)[start]
	}

	pivotIndex := rand.Intn(end-start) + start
	selectedIndex := partition(nums, start, end, pivotIndex)
	if selectedIndex < k {
		return quickSelect(nums, k, selectedIndex+1, end)
	} else if selectedIndex > k {
		return quickSelect(nums, k, start, selectedIndex-1)
	} else {
		return (*nums)[selectedIndex]
	}

}

func main() {
	nums := []int{2, 1, 3, 5, 6, 4}
	k := 3
	fmt.Println(findKthLargest(nums, k))
}
