package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Stack[T any] interface {
	Pop() T
	Push(item T)
	Peek() T
}

type ListNodeStack struct {
	Elements []*ListNode
}

func (s *ListNodeStack) Push(item *ListNode) {
	s.Elements = append(s.Elements, item)
}

func (s *ListNodeStack) Pop() *ListNode {
	lastElement := s.Peek()
	s.Elements = s.Elements[:len(s.Elements)-1]

	return lastElement
}

func (s *ListNodeStack) Peek() *ListNode {
	return s.Elements[len(s.Elements)-1]
}

func doubleIt(head *ListNode) *ListNode {
	stack := ListNodeStack{Elements: make([]*ListNode, 0)}

	for head != nil {
		stack.Push(head)

		head = head.Next
	}

	var currentNode *ListNode = nil
	carry := 0
	for len(stack.Elements) > 0 {
		lastNode := stack.Pop()

		doubledValue := lastNode.Val * 2
		if doubledValue > 9 {
			lastDigit := doubledValue%10 + carry
			carry = 1

			if len(stack.Elements) == 0 {
				mostSignificantDigit := &ListNode{Val: doubledValue / 10, Next: lastNode}
				lastNode.Val = lastDigit
				lastNode.Next = currentNode

				return mostSignificantDigit
			}

			lastNode.Val = lastDigit
		} else {
			lastNode.Val = doubledValue + carry
			carry = 0
		}

		if currentNode != nil {
			lastNode.Next = currentNode
		}

		currentNode = lastNode
	}

	return currentNode
}

func main() {
	//	node2 := ListNode{Val: 9, Next: nil}
	//	node1 := ListNode{Val: 9, Next: &node2}
	//	node0 := ListNode{Val: 9, Next: &node1}

	node0 := ListNode{Val: 5, Next: nil}

	head := doubleIt(&node0)

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
