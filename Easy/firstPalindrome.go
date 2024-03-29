package main

func firstPalindrome(words []string) string {
	for _, val := range words {
		pal := true

		for i := 0; i < len(val)/2; i++ {
			if val[i] != val[len(val)-i-1] {

				pal = false
				break
			}
		}
		if pal == true {
			return val
		}
	}
	return ""

}

// func main() {

// 	fmt.Println(firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
// }
