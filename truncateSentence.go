package main

import (
	"fmt"
	"strings"
)

func truncateSentence(s string, k int) string {
	var result string
	str := strings.Split(s, " ")

	for i, ch := range str {

		if i == k {

			break

		}
		result += ch + " "

	}
	result = strings.TrimSpace(result)
	return result

}
func main() {
	// Примеры использования функции truncateSentence
	s := "Hello how are you Contestant"
	k := 4
	fmt.Println(truncateSentence(s, k)) // Output: "Hello how are you"

	// s2 := "What is the solution to this problem"
	// k2 := 4
	// fmt.Println(truncateSentence(s2, k2)) // Output: "What is the solution"

	// s3 := "chopper is not a tanuki"
	// k3 := 5
	// fmt.Println(truncateSentence(s3, k3)) // Output: "chopper is not a tanuki"
}
