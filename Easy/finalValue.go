package main

func finalValueAfterOperations(operations []string) int {
	var result int
	for _, ch := range operations {
		if ch == "++X" || ch == "X++" {
			result++
		} else {
			result--
		}
	}
	return result

}

// func main() {
// 	fmt.Println(finalValueAfterOperations([]string{"--X","X++","X++"})) //1
// 	fmt.Println(finalValueAfterOperations([]string{"++X","++X","X++"})) //3
// 	fmt.Println(finalValueAfterOperations([]string{"X++","++X","--X","X--"})) //0
// }
