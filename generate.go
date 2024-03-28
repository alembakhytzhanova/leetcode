package main

import "fmt"

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		if i == 0 {
			res = append(res, []int{1})
			continue
		}

		p := []int{}
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				p = append(p, 1)
				continue
			}
			// p = append(p, res[j-1][j]+res[j][j+1])
			fmt.Println(p)
			p = append(p, res[i-1][j-1]+res[i-1][j])
		}
		res = append(res, p)

	}
	return res
}
