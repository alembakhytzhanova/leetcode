package main

import "fmt"

func meanArray(a []int) float64 {
	n := len(a)

	if n == 0 {
		return 0
	}
	sum := sumArray(a, n)
	return float64(sum) / float64(n)

}

func sumArray(a []int, n int) int {
	if n <= 0 {
		return 0
	}
	return sumArray(a, n-1) + a[n-1]
}

func main() {
	arr := []int{1, 2, 3, 4, 5} // Example array
	mean := meanArray(arr)
	fmt.Printf("The mean of the array is: %.2f\n", mean)
}
