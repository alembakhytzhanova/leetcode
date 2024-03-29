package main

func shuffle(nums []int, n int) []int {
	result := []int{}
	for i := 0; i < n; i++ {
		result = append(result, nums[i], nums[i+n])
	}
	return result
}

// func main() {
// 	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
// }
