package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	i := len(students)

	c := len(students) * len(students)

	for i > 0 {
		if sandwiches[0] == sandwiches[0] {
			students = sandwiches[1:]
			sandwiches = sandwiches[1:]
			i--

		} else {
			t := sandwiches[0]
			sandwiches = sandwiches[1:]
			sandwiches = append(sandwiches, t)
			if c == 0 {
				return i
			}
			c--
		}
	}
	return 0
}

func main() {
	fmt.Println(countStudents([]int{1, 1, 0, 0}, []int{0, 1, 0, 1}))
}
