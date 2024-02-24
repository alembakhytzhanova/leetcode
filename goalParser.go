package main

// func interpret(command string) string {
// 	var result string

// 	word := strings.ReplaceAll(command, "()", "o")
// 	result = strings.ReplaceAll(word, "(al)", "al")

// 	return result

// }

func interpret(command string) string {
	var res string
	var comMap = map[string]string{
		"G":    "G",
		"(al)": "al",
		"()":   "o",
	}
	var s string
	for _, v := range command {
		s += string(v)
		if item, ok := comMap[s]; ok {
			res += item
			s = ""
		}
	}
	return res
}
