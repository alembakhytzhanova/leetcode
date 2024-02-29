package main

func findWordsContaining(words []string, x byte) []int {
	var result []int

	for i, ch := range words {
		for _, char := range ch {
			if char == rune(x) {
				result = append(result, i)
				break

			}
		}
	}
	return result
}

// func main() {
// 	fmt.Println(findWordsContaining([]string{"leet", "code"}, byte('e')))              // [0,1]
// 	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('a'))) // [0,2]
// 	fmt.Println(findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, byte('z'))) // []

// }
