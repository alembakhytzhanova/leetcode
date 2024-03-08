package main

import "fmt"

func getConcatenation(nums []int) []int {

	num := len(nums)
	m := make([]int, 2*num)

	for i := 0; i < num; i++ {
		fmt.Println("this is i:", i)
		m[i], m[i+num] = nums[i], nums[i]

	}
	return m
}

// func main() {
// 	fmt.Println(getConcatenation([]int{1, 2, 1}))
// }
