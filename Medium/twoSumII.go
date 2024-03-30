package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	var result []int
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			result = append(result, l+1, r+1)
			return result

		} else if sum < target {
			l++
		} else {
			r--
		}

	}

	return result

}
func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
