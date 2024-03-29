package main

import "fmt"

func restoreString(s string, indices []int) string {
	var result string
	for _, letter := range s {
		fmt.Println(letter)
		for _, num := range indices {
			fmt.Println(num)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3})) //leetcode
// 	fmt.Println(restoreString("abc", []int{0, 1, 2}))                     //abc
// }

// func restoreString(s string, indices []int) string {
// 	result := make([]byte, len(s))
// 	for l,index := range indices {
// 		result[index] = s[l]
// 	}
// 	return string(result)
// }
