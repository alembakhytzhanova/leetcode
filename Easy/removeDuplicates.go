package main

func removeDuplicates(s string) string {
	if len(s) <= 0 {
		return "0"
	}
	stack := []rune{}

	for _, ch := range s {
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			//fmt.Println("len(stack)", stack[len(stack)-1])

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
			//fmt.Println("stack", "ch", stack, ch)

		}
	}
	return string(stack)
}

// func main() {
// 	fmt.Println(removeDuplicates("abbaca"))
// 	fmt.Println(removeDuplicates("azxxzy"))
// }
