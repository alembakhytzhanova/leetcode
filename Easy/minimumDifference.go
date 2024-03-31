package main

import (
	"fmt"
	"sort"
)

// func minimumDifference(nums []int, k int) int {
// 	sort.Ints(nums)
// 	min:= 0
// 	max:= 0
// 	fmt.Println("sort", nums)
// 	res := 0
// 	l, r := 0, len(nums)-1
// 	for l < r {
// 		res = nums[r] - nums[l]
// 		fmt.Println("res", res)

//		}
//		return res
//	}
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, k-1
	minRes := nums[r] - nums[l]
	for r < len(nums) {
		minRes = minimal(minRes, nums[r]-nums[l])
		l, r = l+1, r+1

	}
	return minRes
}

func minimal(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func main() {
	//fmt.Println(minimumDifference([]int{87063, 61094, 44530, 21297, 95857, 93551, 9918}, 6))
	fmt.Println(minimumDifference([]int{9, 4, 1, 7}, 2))
}
