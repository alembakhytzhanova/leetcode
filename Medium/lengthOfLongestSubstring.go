package main

func lengthOfLongestSubstring(s string) int {
	// Создаем map для отслеживания индекса каждого символа в строке
	charMap := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		// Если текущий символ уже встречался в текущем подстроке,
		// обновляем начало подстроки
		if index, found := charMap[s[end]]; found && index >= start {
			start = index + 1
		}
		// Обновляем индекс текущего символа
		charMap[s[end]] = end
		// Обновляем максимальную длину подстроки
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

// func main() {
// 	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // Output: 3
// 	fmt.Println(lengthOfLongestSubstring("bbbbb"))    // Output: 1
// 	fmt.Println(lengthOfLongestSubstring("pwwkew"))   // Output: 3
// }
