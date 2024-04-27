package main

import "log"

type Stack struct {
	content []int
}

//func main() {
	s := Stack{}
	s.Push(12)
	s.Push(8)
	s.Display()

	log.Print(s.Pop())
	log.Print(s.Pop())
	log.Print(s.Pop())

	s.Display()
}

func (stack *Stack) Display() {
	log.Print(stack.content)
}

func (stack *Stack) Push(item int) {
	stack.content = append(stack.content, item)
}

func (stack *Stack) Pop() int {
	if len(stack.content) == 0 {
		return 0
	}
	lastItemIndex := len(stack.content) - 1
	lastItem := stack.content[lastItemIndex]

	stack.content = stack.content[:lastItemIndex]
	return lastItem

}
