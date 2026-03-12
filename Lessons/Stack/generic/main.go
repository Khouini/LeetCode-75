package main

import (
	"errors"
	"fmt"
)

// Stack is a generic LIFO data structure
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty returns true if the stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Clear removes all elements from the stack
func (s *Stack[T]) Clear() {
	s.items = nil
}

func main() {
	// --- Integer Stack ---
	fmt.Println("=== Integer Stack ===")
	var intStack Stack[int]

	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println("Size:", intStack.Size()) // 3

	if top, err := intStack.Peek(); err == nil {
		fmt.Println("Peek:", top) // 30
	}

	for !intStack.IsEmpty() {
		val, _ := intStack.Pop()
		fmt.Println("Popped:", val) // 30, 20, 10
	}

	_, err := intStack.Pop()
	fmt.Println("Pop from empty:", err) // stack is empty

	// --- String Stack ---
	fmt.Println("\n=== String Stack ===")
	var strStack Stack[string]

	strStack.Push("go")
	strStack.Push("is")
	strStack.Push("awesome")

	fmt.Println("Size:", strStack.Size()) // 3
	strStack.Clear()
	fmt.Println("Size after Clear:", strStack.Size()) // 0

	// --- Balanced Parentheses using Stack ---
	fmt.Println("\n=== Balanced Parentheses ===")
	tests := []string{"(())", "({[]})", "([)]", "(("}
	for _, s := range tests {
		fmt.Printf("%-10s → %v\n", s, isBalanced(s))
	}
}

// isBalanced checks if brackets in a string are balanced
func isBalanced(s string) bool {
	var stack Stack[rune]
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, ch := range s {
		switch ch {
		case '(', '[', '{':
			stack.Push(ch)
		case ')', ']', '}':
			top, err := stack.Pop()
			if err != nil || top != pairs[ch] {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
