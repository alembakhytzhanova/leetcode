package main

import "fmt"

func printN(n int) {
	if n > 0 {
		printN(n - 1)
		fmt.Println(n)
	}
}

// func main() {
// 	n := 10
// 	printN(n)
// }
