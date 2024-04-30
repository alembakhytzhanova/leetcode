package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var a, b int
		fmt.Scan(&a, &b)

		sum := a + b

		fmt.Println(sum)
	}
}

// 1 задание
