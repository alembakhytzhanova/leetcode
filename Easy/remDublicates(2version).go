package main

// func main() {
// 	s := "azxxzy"
// 	fmt.Println(Remove(s))

// }

type Stack struct {
	data []string
}

func Remove(s string) string {
	var result string
	stack := Stack{}

	for _, ch := range s {
		if !stack.isEmpty() && string(ch) == stack.Peek() {
			stack.Pop()

		} else {
			stack.Push(string(ch))
		}
	}
	for !stack.isEmpty() {
		result = stack.Pop() + result
	}
	return result
}

func (s *Stack) Peek() string {
	item := s.data[len(s.data)-1]
	return item

}

func (s *Stack) Pop() string {
	item := s.Peek()
	s.data = s.data[:len(s.data)-1]
	return item
}

func (s *Stack) Push(item string) {

	s.data = append(s.data, item)
}

func (s *Stack) isEmpty() bool {
	if len(s.data) == 0 {
		return true
	}
	return false
}
