package main

import "fmt"

// input: []int{1, 2 , 3}, []int{2, 3, 4}, []int{3, 4, 5} if nums are unique
// output: 共同有數字 且為最大值

func findMaxInArrays(nums1, nums2, nums3 []int) int {

	m := make(map[int]int)

	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		m[v]++
	}

	for _, v := range nums3 {
		m[v]++
	}

	// m = [1 => 1, 2 => 2, 3 => 3, 4 => 2, 5 => 1]
	// find the maximum value in the map
	maxKey := -1
	tempVal := -1
	for k, v := range m {
		if v > tempVal && v >= 3 {
			maxKey = k
		}
	}

	return maxKey
}

func main() {
	res := findMaxInArrays([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
	fmt.Println(res)
}
