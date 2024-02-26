package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	str1 := strings.Join(word1, "")
	str2 := strings.Join(word2, "")
	if str1 == str2 {
		return true
	}
	return false
}

func main() {
	fmt.Println(arrayStringsAreEqual([]string{"ab", "c"}, []string{"a", "bc"}))           //true
	fmt.Println(arrayStringsAreEqual([]string{"a", "cb"}, []string{"ab", "c"}))           //false
	fmt.Println(arrayStringsAreEqual([]string{"abc", "d", "defg"}, []string{"abcddefg"})) //true
}
