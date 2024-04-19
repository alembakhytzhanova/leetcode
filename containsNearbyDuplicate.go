package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	l := 0
	r := 0

	for i := 0; i < k; i++ {
		l += nums[i]
	}
	r = l

	for i := k; i < len(nums); i++ {
		fmt.Println(l, r)
		r = r - nums[i-k] + nums[i]
		l = max(l, r)

	}
	return false
}

func max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}

// func main() {
// 	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
// }
