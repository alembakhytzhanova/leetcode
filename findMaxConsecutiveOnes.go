package main

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0

	for _, ch := range nums {
		if ch == 1 {
			count++
			if count > maxCount {
				maxCount = count

			}

		} else {

			count = 0

		}

	}
	if count > maxCount {
		maxCount = count
		//fmt.Println("count3", count)
		//fmt.Println("MAXcount3", maxCount)
	}
	return maxCount
}

// func main() {
// 	//fmt.Println(findMaxConsecutiveOnes([]int{1, 1})) // 3

// 	fmt.Println(findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})) // 3
// 	//	fmt.Println(findMaxConsecutiveOnes([]int{1, 0, 1, 1, 0, 1})) // 2

// }
