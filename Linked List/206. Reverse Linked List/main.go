package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	realNext := head.Next
	lastProcessedNode := head

	iterator := realNext
	for iterator != nil {
		realNext = iterator.Next
		iterator.Next = lastProcessedNode
		lastProcessedNode = iterator
		if realNext == nil {
			head.Next = nil
			head = iterator
			break
		}
		iterator = realNext
	}

	return head
}

func main() {

	// 1. empty list
	var empty *ListNode

	// 2. single node [1]
	single := &ListNode{Val: 1}

	// 3. two nodes [1 -> 2]
	two := &ListNode{Val: 1}
	two.Next = &ListNode{Val: 2}

	// 4. normal case [1 -> 2 -> 3 -> 4 -> 5]
	five := &ListNode{Val: 1}
	five.Next = &ListNode{Val: 2}
	five.Next.Next = &ListNode{Val: 3}
	five.Next.Next.Next = &ListNode{Val: 4}
	five.Next.Next.Next.Next = &ListNode{Val: 5}

	// run all cases
	runCase("Empty list", empty)
	runCase("Single node", single)
	runCase("Two nodes", two)
	runCase("Five nodes", five)
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func runCase(name string, head *ListNode) {
	fmt.Println("\n---", name, "---")
	fmt.Print("Original: ")
	printList(head)

	head = reverseList(head)

	fmt.Print("Reversed: ")
	printList(head)
}
