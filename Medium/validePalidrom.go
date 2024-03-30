package main

import (
	"strings"
)

func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	l, r := 0, len(s)-1
	for l < r {
		if !isAlphanumeric(s[l]) {
			l++
			continue

		} else if !isAlphanumeric(s[r]) {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--

	}
	return true
}
func isAlphanumeric(char byte) bool {
	return (char >= '0' && char <= '9') || (char >= 'a' && char <= 'z')
}

// func main() {
// 	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
// 	fmt.Println(isPalindrome("race a car"))
// 	fmt.Println(isPalindrome(" "))

// }
