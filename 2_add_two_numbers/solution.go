package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type Nodes struct {
	Node1   *ListNode
	Node2   *ListNode
	DidSwap bool
}

func findNodes(l1 *ListNode, l2 *ListNode) *Nodes {
	node1 := l1
	node2 := l2

	length1 := 0
	length2 := 0

	for node1 != nil {
		length1 += 1
		node1 = node1.Next
	}

	for node2 != nil {
		length2 += 1
		node2 = node2.Next
	}

	if length1 > length2 {
		return &Nodes{
			Node1:   l1,
			Node2:   l2,
			DidSwap: false,
		}
	}

	return &Nodes{
		Node1:   l2,
		Node2:   l1,
		DidSwap: true,
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nodes := findNodes(l1, l2)
	node1 := nodes.Node1
	node2 := nodes.Node2

	carry := 0
	for node1 != nil {
		if node2 == nil {
			node1.Val += carry
			carry = 0
			if node1.Val > 9 {
				carry = 1
				node1.Val = node1.Val % 10

				if node1.Next == nil {
					newTail := &ListNode{Val: 1, Next: nil}
					node1.Next = newTail

					if nodes.DidSwap {
						return l2
					}

					return l1
				}
			}
			node1 = node1.Next

			continue
		}

		node1.Val += node2.Val + carry
		carry = 0

		if node1.Val > 9 {
			carry = 1
			node1.Val = node1.Val % 10

			if node1.Next == nil {
				newTail := &ListNode{Val: 1, Next: nil}
				node1.Next = newTail

				if nodes.DidSwap {
					return l2
				}

				return l1
			}
		}

		node1 = node1.Next
		node2 = node2.Next
	}

	if nodes.DidSwap {
		return l2
	}

	return l1
}

func main() {
	// List1
	//	node06 := ListNode{Val: 9, Next: nil}
	//	node05 := ListNode{Val: 9, Next: &node06}
	//	node04 := ListNode{Val: 9, Next: &node05}
	//	node03 := ListNode{Val: 9, Next: &node04}
	//	node02 := ListNode{Val: 9, Next: &node03}
	//	node01 := ListNode{Val: 9, Next: &node02}
	//	node00 := ListNode{Val: 9, Next: &node01}
	//
	//	// List2
	//	node13 := ListNode{Val: 9, Next: nil}
	//	node12 := ListNode{Val: 9, Next: &node13}
	//	node11 := ListNode{Val: 9, Next: &node12}
	//	node10 := ListNode{Val: 9, Next: &node11}
	//

	node02 := ListNode{Val: 9, Next: nil}
	node01 := ListNode{Val: 4, Next: &node02}
	node00 := ListNode{Val: 2, Next: &node01}

	// List2
	node13 := ListNode{Val: 9, Next: nil}
	node12 := ListNode{Val: 4, Next: &node13}
	node11 := ListNode{Val: 6, Next: &node12}
	node10 := ListNode{Val: 5, Next: &node11}
	head := addTwoNumbers(&node00, &node10)

	for head != nil {
		println(head.Val)

		head = head.Next
	}
}
