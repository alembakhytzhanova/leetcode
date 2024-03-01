package main

func cellsInRange(s string) []string {
	var result []string
	var a, b byte = s[0], s[3]

	var c, d byte = s[1], s[4]

	for i := a; i <= b; i++ {

		for j := c; j <= d; j++ {

			result = append(result, string(i)+string(j))
		}

	}
	return result
}

// func main() {
// 	fmt.Println(cellsInRange("K1:L2"))
// }
