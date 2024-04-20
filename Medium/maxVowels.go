package main

import "fmt"

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	currentVowel := 0
	vowelMax := 0

	for i := 0; i < len(s); i++ {
		if i >= k {
			// if len > k , discard left symbol
			if vowels[s[i-k]] {

				vowelMax--
			}
		}
		//check vowel a new symbol
		if vowels[s[i]] {
			vowelMax++

		}

		if vowelMax > currentVowel {
			currentVowel = vowelMax
		}
	}
	return currentVowel

}

func main() {
	fmt.Println(maxVowels("abciiidef", 3))
	fmt.Println(maxVowels("leetcode", 3))
}
