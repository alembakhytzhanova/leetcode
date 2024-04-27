package main

import "fmt"

func removeStars(s string) string {
	stack := []rune{}
	for _, ch := range s {
		if len(stack) > 0 && ch == '*' {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}

func main() {
	fmt.Println(removeStars("leet**cod*e"))
}
