package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = []int{}

func inorderThisThing(root *TreeNode) {
	if root == nil {
		return
	}

	inorderThisThing(root.Left)

	res = append(res, root.Val)

	inorderThisThing(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	inorderThisThing(root)
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

	result := inorderTraversal(root)

	for _, value := range result {
		println(value)
	}
}
