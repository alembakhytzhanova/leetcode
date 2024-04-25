package main

func findAnagrams(s string, p string) []int {
	mp := map[byte]int{}

	k := len(p)
	r := []int{}
	for i := 0; i < len(p); i++ {
		mp[p[i]]++
	}

	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++

		if i >= k-1 {
			if isAnagram(mp, m) {
				r = append(r, i-k+1)
			}
			m[s[i-k+1]]--
		}
	}
	return r
}

func isAnagram(m1, m2 map[byte]int) bool {
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
