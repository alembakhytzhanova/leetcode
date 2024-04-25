package main

import "fmt"

func countStudents(students []int, sandwiches []int) int {
	i := len(students)
	c := len(students) * len(students)

	for i > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
			i--
		} else {
			t := students[0]
			students = students[1:]
			students = append(students, t)
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
