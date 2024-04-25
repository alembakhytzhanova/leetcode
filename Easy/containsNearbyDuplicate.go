package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	l := 0
	for r, num := range nums {
		if r-l > k {
			delete(m, nums[l])
			l++
		}
		if _, found := m[num]; found {
			return true
		}
		m[num] = r

	}
	return false
}
func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}
