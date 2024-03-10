package main

import "fmt"

func sumOfMultiples(n int) int {
	sum := 0

	for i := 0; i <= n; i++ {
		if i%3 == 0 || i%5 == 0 || i%7 == 0 {
			sum += i
		}
	}
	return sum
}

func main() {
	fmt.Println(sumOfMultiples(10))
}
