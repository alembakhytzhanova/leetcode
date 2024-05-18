package main

import "fmt"

func printNto1(n int) {
	if n > 0 {
		fmt.Println(n)
		printNto1(n - 1)

	}
}

// func main() {
// 	n := 10
// 	printNto1(n)
// }
