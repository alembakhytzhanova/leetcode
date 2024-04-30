package main

import (
	"fmt"
	"sort"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)

		times := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&times[j])
		}

		rankings := make(map[int]int) // Времена и соответствующие им места
		sort.Ints(times)              // Сортируем времена

		for j, time := range times {
			if _, ok := rankings[time]; !ok { // Если время не присутствует в словаре, присваиваем место
				rankings[time] = j + 1
			}
		}

		// Выводим места спортсменов
		for _, time := range times {
			fmt.Printf("%d ", rankings[time])
		}
		fmt.Println()
	}
}
