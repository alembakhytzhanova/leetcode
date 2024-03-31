package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := 0

	for i := 1; i < len(nums); i++ {
		if nums[l] != nums[i] {
			l++

			nums[l] = nums[i]

		}

	}
	nums = nums[:l+1]

	return len(nums)

}

func main() {
	//fmt.Println(removeDuplicates([]int{1, 1, 2}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}
