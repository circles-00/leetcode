package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val = node.Next.Val

	node.Next = node.Next.Next
}

func main() {
	Node3 := ListNode{Val: 9, Next: nil}
	Node2 := ListNode{Val: 1, Next: &Node3}
	Node1 := ListNode{Val: 5, Next: &Node2}
	Node0 := ListNode{Val: 4, Next: &Node1}

	deleteNode(&Node1)

	fmt.Print("[")

	fmt.Printf("%d,", Node0.Val)
	fmt.Printf("%d,", Node1.Val)
	fmt.Printf("%d,", Node2.Val)
	fmt.Printf("%d", Node3.Val)

	fmt.Print("]")
}
