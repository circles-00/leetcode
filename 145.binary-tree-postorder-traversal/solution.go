package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = []int{}

func postorderThisThing(root *TreeNode) {
	if root == nil {
		return
	}

	postorderThisThing(root.Left)
	postorderThisThing(root.Right)

	res = append(res, root.Val)
}

func postorderTraversal(root *TreeNode) []int {
	postorderThisThing(root)
	newRes := make([]int, len(res))
	copy(newRes, res)
	res = []int{}

	return newRes
}

func main() {
	// [1,null,2,3]

	right1 := &TreeNode{
		Val: 3,
	}

	right0 := &TreeNode{
		Val:  2,
		Left: right1,
	}

	root := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: right0,
	}

	result := postorderTraversal(root)

	for _, value := range result {
		println(value)
	}
}
