package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	var result int

	for _, word := range sentences {
		newSentence := strings.Split(word, " ")
		fmt.Println(newSentence)

		lenMaxSentence := len(newSentence)
		if lenMaxSentence > result {
			result = lenMaxSentence
		}

	}

	return result

}

// func main() {
// 	case1 := mostWordsFound([]string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}) //6
// 	case2 := mostWordsFound([]string{"please wait", "continue to fight", "continue to win"})                             //3
// 	fmt.Println(case1)
// 	fmt.Println(case2)

// }
