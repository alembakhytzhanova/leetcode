package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s) 

	var n int
	fmt.Scanln(&n) 

	result := s 

	for i := 0; i < n; i++ {
		var start, end int
		var r string
		fmt.Scanln(&start, &end, &r) 

		result = replaceSubstring(result, start-1, end-1, r)
	}

	fmt.Println(result) 
}


func replaceSubstring(s string, start, end int, r string) string {
	return s[:start] + r + s[end+1:]
}
 // 2 zadanie