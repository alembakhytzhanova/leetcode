package main

import "fmt"

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0
	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1]
            j++
		}
	}
    fmt.Println(stack)
	return len(stack) == 0
}


func main() {

}