package main

import (
	"fmt"
	"math"
)

func minEatingSpeed(piles []int, h int) int {
	l, r := 1, maximum(piles)

	res := 0
	for l <= r {
		k := (l + r) / 2
		hours := 0
		for _, v := range piles {

			hours += int(math.Ceil(float64(v) / float64(k)))

		}
		if hours <= h {
			res = k
			r = k - 1
		} else {
			l = k + 1

		}
	}

	return res
}

func maximum(m []int) int {
	maxInt := 0
	for _, ch := range m {
		if maxInt < ch {
			maxInt = ch
		}
	}
	return maxInt
}

// func minimum(m []int) int {
// 	minInt := int(1e9)
// 	for _, ch := range m {
// 		if minInt > ch {
// 			minInt = ch
// 		}
// 	}
// 	return minInt
// }

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
}
