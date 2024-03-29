package main

func numJewelsInStones(jewels string, stones string) int {
	result := 0
	for _, ch := range jewels {
		for _, v := range stones {
			if ch == v {
				result++
			}
		}

	}
	return result
}
