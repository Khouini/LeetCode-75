package main

import "fmt"

// ─────────────────────────────────────────────
//  Node & LinkedList definitions
// ─────────────────────────────────────────────

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
	Size int
}

// ─────────────────────────────────────────────
//  Insert Operations
// ─────────────────────────────────────────────

// InsertAtHead inserts a new node at the beginning — O(1)
func (l *LinkedList) InsertAtHead(val int) {
	newNode := &Node{Value: val}
	newNode.Next = l.Head
	l.Head = newNode
	l.Size++
}

// InsertAtTail inserts a new node at the end — O(n)
func (l *LinkedList) InsertAtTail(val int) {
	newNode := &Node{Value: val}

	if l.Head == nil {
		l.Head = newNode
		l.Size++
		return
	}

	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	l.Size++
}

// InsertAt inserts a new node at a given index — O(n)
func (l *LinkedList) InsertAt(index, val int) bool {
	if index < 0 || index > l.Size {
		fmt.Printf("Index %d out of bounds (size: %d)\n", index, l.Size)
		return false
	}

	if index == 0 {
		l.InsertAtHead(val)
		return true
	}

	newNode := &Node{Value: val}
	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	newNode.Next = current.Next
	current.Next = newNode
	l.Size++
	return true
}

// ─────────────────────────────────────────────
//  Delete Operations
// ─────────────────────────────────────────────

// DeleteByValue removes the first node with the given value — O(n)
func (l *LinkedList) DeleteByValue(val int) bool {
	if l.Head == nil {
		return false
	}

	if l.Head.Value == val {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	current := l.Head
	for current.Next != nil {
		if current.Next.Value == val {
			current.Next = current.Next.Next
			l.Size--
			return true
		}
		current = current.Next
	}
	return false
}

// DeleteAt removes the node at the given index — O(n)
func (l *LinkedList) DeleteAt(index int) bool {
	if index < 0 || index >= l.Size || l.Head == nil {
		fmt.Printf("Index %d out of bounds (size: %d)\n", index, l.Size)
		return false
	}

	if index == 0 {
		l.Head = l.Head.Next
		l.Size--
		return true
	}

	current := l.Head
	for i := 0; i < index-1; i++ {
		current = current.Next
	}
	current.Next = current.Next.Next
	l.Size--
	return true
}

// ─────────────────────────────────────────────
//  Search & Access
// ─────────────────────────────────────────────

// Search returns true if the value exists in the list — O(n)
func (l *LinkedList) Search(val int) bool {
	current := l.Head
	for current != nil {
		if current.Value == val {
			return true
		}
		current = current.Next
	}
	return false
}

// GetAt returns the value at a given index — O(n)
func (l *LinkedList) GetAt(index int) (int, bool) {
	if index < 0 || index >= l.Size {
		return 0, false
	}

	current := l.Head
	for i := 0; i < index; i++ {
		current = current.Next
	}
	return current.Value, true
}

// ─────────────────────────────────────────────
//  Utility Operations
// ─────────────────────────────────────────────

// Print displays the list in a readable format
func (l *LinkedList) Print() {
	if l.Head == nil {
		fmt.Println("(empty list)")
		return
	}
	current := l.Head
	for current != nil {
		fmt.Printf("%d", current.Value)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Println(" -> nil")
}

// Reverse reverses the linked list in place — O(n)
func (l *LinkedList) Reverse() {
	var prev *Node
	current := l.Head

	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}

// Length returns the number of nodes
func (l *LinkedList) Length() int {
	return l.Size
}

// ToSlice converts the linked list to a Go slice
func (l *LinkedList) ToSlice() []int {
	result := make([]int, 0, l.Size)
	current := l.Head
	for current != nil {
		result = append(result, current.Value)
		current = current.Next
	}
	return result
}

// HasCycle detects if the list has a cycle (Floyd's algorithm) — O(n)
func (l *LinkedList) HasCycle() bool {
	slow := l.Head
	fast := l.Head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// ─────────────────────────────────────────────
//  Main — Demo
// ─────────────────────────────────────────────

func main() {
	fmt.Println("════════════════════════════════")
	fmt.Println("       Linked List in Go        ")
	fmt.Println("════════════════════════════════")

	list := &LinkedList{}

	// Insert operations
	fmt.Println("\n── Insertions ──")
	list.InsertAtTail(10)
	list.InsertAtTail(20)
	list.InsertAtTail(30)
	list.InsertAtTail(40)
	list.InsertAtHead(5)
	list.Print() // 5 -> 10 -> 20 -> 30 -> 40 -> nil

	list.InsertAt(2, 15) // insert 15 at index 2
	fmt.Print("After InsertAt(2, 15): ")
	list.Print() // 5 -> 10 -> 15 -> 20 -> 30 -> 40 -> nil

	// Size
	fmt.Printf("\nList size: %d\n", list.Length())

	// Search
	fmt.Println("\n── Search ──")
	fmt.Printf("Search(20): %v\n", list.Search(20)) // true
	fmt.Printf("Search(99): %v\n", list.Search(99)) // false

	// Access by index
	fmt.Println("\n── GetAt ──")
	if val, ok := list.GetAt(3); ok {
		fmt.Printf("GetAt(3): %d\n", val) // 20
	}

	// Delete operations
	fmt.Println("\n── Deletions ──")
	list.DeleteByValue(15)
	fmt.Print("After DeleteByValue(15): ")
	list.Print()

	list.DeleteAt(0)
	fmt.Print("After DeleteAt(0):       ")
	list.Print()

	list.DeleteAt(3)
	fmt.Print("After DeleteAt(3):       ")
	list.Print()

	// Reverse
	fmt.Println("\n── Reverse ──")
	fmt.Print("Before: ")
	list.Print()
	list.Reverse()
	fmt.Print("After:  ")
	list.Print()

	// Convert to slice
	fmt.Println("\n── ToSlice ──")
	fmt.Printf("As slice: %v\n", list.ToSlice())

	// Cycle detection
	fmt.Println("\n── Cycle Detection ──")
	fmt.Printf("Has cycle: %v\n", list.HasCycle()) // false

	fmt.Println("\n════════════════════════════════")
}
