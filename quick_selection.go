// Example (LeetCode 215) Kth Largest Element in an Array

package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	// if there are 6 numbers in nums, the 2nd largest = 5th (6-2+1)th smallest element.
	// nums[4] is 5th smallest element.
	return quickSelect(&nums, len(nums)-k, 0, len(nums)-1)
}

// ref https://shubo.io/quick-sort/#partition-%E5%A6%82%E4%BD%95%E9%81%8B%E4%BD%9C
func partition(nums *[]int, start, end, pivotIndex int) int {
	pivotValue := (*nums)[pivotIndex]

	// Move pivot to the end (the end value is seen as pivot)
	(*nums)[end], (*nums)[pivotIndex] = (*nums)[pivotIndex], (*nums)[end]

	// Move smaller values to left of the pivotValue
	// 用 nextLeftIndex 紀錄下一個小於 pivot 的元素要交換到的位置
	nextLeftIndex := start
	for i := start; i < end; i++ {
		if (*nums)[i] < pivotValue {
			(*nums)[nextLeftIndex], (*nums)[i] = (*nums)[i], (*nums)[nextLeftIndex]
			// 最後 nextLeftIndex 會指到一個大於 pivot 的值
			nextLeftIndex++
		}
	}

	// 此時 pivot 在 end 位置上，將 pivot 與 nextLeftIndex 所指的值交換，如此就完成 partition
	(*nums)[nextLeftIndex], (*nums)[end] = (*nums)[end], (*nums)[nextLeftIndex]

	// 回傳 pivot 經過一次 partition 後的位置
	return nextLeftIndex
}

func quickSelect(nums *[]int, k, start, end int) int {

	if start == end {
		return (*nums)[start]
	}

	// select a random pivot within [start, ... ,end]
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
