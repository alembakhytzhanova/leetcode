package main

import (
	"fmt"
	"strings"
	"unicode"
)

func SortSentence(s string) string {
	str1 := strings.Split(s, " ")
	m := map[string]string{}
	ints := "123456789"
	result := ""
	for _, v := range str1 {
		for _, h := range v {
			for _, v2 := range ints {
				if h == v2 {
					m[string(v2)] = v

					break
				}
			}
		}
	}

	for i, v := range ints {
		if m[string(v)] == "" {

			break
		}
		result += m[string(v)]

		if i != len(str1)-1 {
			result += " "
		}
	}
	res := ""
	for _, ch := range result {
		if !unicode.IsDigit(ch) {
			res += string(ch)
		}
	}
	return res
}

func main() {
	fmt.Println(SortSentence("is2 sentence4 This1 a3"))
}
