package main

type ArrayQueue[T any] struct {
	Elements []T
	Size     int
}

func (q *ArrayQueue[T]) Enqueue(item T) {
	q.Elements = append(q.Elements, item)
	q.Size += 1
}

func (q *ArrayQueue[T]) Dequeue() T {
	frontElement := q.Peek()

	q.Elements = q.Elements[1:]
	q.Size -= 1

	return frontElement
}

func (q *ArrayQueue[T]) Peek() T {
	if len(q.Elements) > 0 {
		return q.Elements[0]
	}

	var empty_result T

	return empty_result
}

func (q *ArrayQueue[T]) Clear() {
	q.Elements = make([]T, 0)
	q.Size = 0
}

func (q *ArrayQueue[T]) IsEmpty() bool {
	return q.Size == 0
}

func NewArrayQueue[T any]() *ArrayQueue[T] {
	return &ArrayQueue[T]{Elements: make([]T, 0), Size: 0}
}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	visited := make(map[int]bool)
	queue := NewArrayQueue[*Node]()

	visited[root.Val] = true
	queue.Enqueue(root)

	result := [][]int{}

	for !queue.IsEmpty() {
		levelSize := queue.Size
		currentLevel := []int{}

		for i := 0; i < levelSize; i++ {
			node := queue.Dequeue()
			currentLevel = append(currentLevel, node.Val)

			if node.Children != nil {
				for _, child := range node.Children {
					queue.Enqueue(child)
				}
			}
		}

		result = append(result, currentLevel)
	}

	return result
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

	result := levelOrder(root)

	for i, value := range result {
		for _, v := range value {
			println("index:", i, "value:", v)
		}
	}
}
