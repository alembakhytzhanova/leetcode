package main

import "fmt"

func maxWidthRamp(nums []int) int {
	var stack []int
	var res int

	for i := 0; i < len(nums); i++ {
		if len(stack) == 0 || nums[i] < nums[stack[len(stack)-1]] {
			stack = append(stack, i)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] >= nums[stack[len(stack)-1]] {
			width := i - stack[len(stack)-1]
			if width > res {
				res = width
			}
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)

	}

	return res
}

func main() {
	fmt.Println(maxWidthRamp([]int{6, 0, 8, 2, 1, 5})) //4
}
