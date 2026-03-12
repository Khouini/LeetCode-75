package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	last := len(s.items) - 1
	top := s.items[last]
	s.items = s.items[:last]
	return top, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func main() {
	var s Stack

	s.Push(10)
	s.Push(20)
	s.Push(30)

	fmt.Println("Size:", s.Size())

	if top, err := s.Peek(); err == nil {
		fmt.Println("Peek:", top)
	}

	for !s.IsEmpty() {
		val, _ := s.Pop()
		fmt.Println("Popped:", val)
	}

	_, err := s.Pop()
	fmt.Println("Pop from empty:", err)
}
