package main

import (
	"strings"
)

func reverseWords(s string) string {
	var result string
	var word string
	newStr := strings.Split(s, "")
	for _, ch := range newStr {
		if ch != " " {
			word = ch + word
		} else {
			result += word + ch
			word = ""
		}
	}
	return result + word

}

// func main() {
// 	fmt.Println(reverseWords("Let's take LeetCode contest"))
// }
https://leetcode.com/problems/reverse-words-in-a-string-iii/submissions/1180022653/