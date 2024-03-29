package main

func countConsistentStrings(allowed string, words []string) int {

	counter := 0
	for _, word := range words {
		for _, ch := range word {
			if !contain(allowed, ch) {
				counter--
				break
			}
		}
		counter++
	}
	return counter
}

func contain(s string, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
