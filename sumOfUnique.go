package main

import "fmt"

func sumOfUnique(nums []int) int {

	occurrences := make(map[int]int)

	for _, num := range nums {
		occurrences[num]++
	}

	sum := 0

	for num, count := range occurrences {
		if count == 1 {
			sum += num
		}
	}

	return sum
}

func main() {
	fmt.Println(sumOfMultiples([]int{1, 2, 3, 2}))
}
