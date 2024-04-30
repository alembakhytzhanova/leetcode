package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res, stack := make([]int, n), []int{}
	for i := n - 1; i >= 0; i-- {

		for len(stack) != 0 && temperatures[i] >= temperatures[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) != 0 {
			res[i] = stack[len(stack)-1] - i
		}

		stack = append(stack, i)
	}
	return res
}
