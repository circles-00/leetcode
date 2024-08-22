package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}
	curr := result

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next
	}

	if list1 == nil {
		curr.Next = list2
	} else {
		curr.Next = list1
	}

	return result.Next
}

func main() {
	node02 := &ListNode{Val: 4}
	node01 := &ListNode{Val: 2, Next: node02}
	node00 := &ListNode{Val: 1, Next: node01}

	node12 := &ListNode{Val: 4}
	node11 := &ListNode{Val: 3, Next: node12}
	node10 := &ListNode{Val: 1, Next: node11}

	result := mergeTwoLists(node00, node10)

	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
