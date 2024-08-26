package main

type Node struct {
	Val      int
	Children []*Node
}

var res = []int{}

func postorderThisThing(root *Node) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		postorderThisThing(child)
	}

	res = append(res, root.Val)
}

func postorderTraversal(root *Node) []int {
	postorderThisThing(root)
	newRes := make([]int, len(res))
	copy(newRes, res)
	res = []int{}

	return newRes
}

func postorder(root *Node) []int {
	postorderThisThing(root)
	newRes := make([]int, len(res))
	copy(newRes, res)
	res = []int{}

	return newRes
}

func main() {
	// [1,null,3,2,4,null,5,6]

	child2 := &Node{
		Val: 4,
	}

	child1 := &Node{
		Val: 2,
	}

	child01 := &Node{
		Val: 5,
	}

	child02 := &Node{
		Val: 6,
	}

	child0 := &Node{
		Val:      3,
		Children: []*Node{child01, child02},
	}

	root := &Node{
		Val:      1,
		Children: []*Node{child0, child1, child2},
	}

	result := postorder(root)

	for _, value := range result {
		println(value)
	}
}
