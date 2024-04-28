package main 



func isValid(s string) bool {
	stack := []rune{}
    
    for _ , v := range s {
        if len(stack) >= 3 && string(stack[len(stack)-3:]) == "abc" {
            stack = stack[:len(stack)-3]
        }
        stack = append(stack,v)
    }

    return "abc" == string(stack)

}