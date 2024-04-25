package main

import "fmt"

func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 != 0 {
		return false
	}
	m := map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := []rune{}

	for _, ch := range s {
		if _, ok := m[ch]; ok {
			stack = append(stack, ch)
			continue
		}
		if len(stack) == 0 || m[stack[len(stack)-1]] != ch {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
