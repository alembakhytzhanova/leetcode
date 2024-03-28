package main

func sortedSquares(nums []int) []int {
	var result []int

	for _, ch := range nums {
		sq := ch * ch
		result = append(result, sq)
	}
	for i := 0; i < len(result)-1; i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}

		}

	}
	return result
}
