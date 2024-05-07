package main

type Queue interface {
  Enqueue(item int)
  Dequeue()
  Peek() int
}

type SliceQueue struct {
  Elements []int
  Size int
}

func (q * SliceQueue) Enqueue(item int) {
  q.Elements = append(q.Elements, item)
  q.Size += 1
}

func (q * SliceQueue) Dequeue() {
  q.Elements = q.Elements[1:]
  q.Size -= 1
}

// Returns -1 if the queue is empty
func (q SliceQueue) Peek() int {
  if len(q.Elements) > 0 {
    return q.Elements[0]
  }

  return -1
}

func findMinHeightTrees(n int, edges [][]int) []int {
  visited := make(map[int]bool)
  queue := SliceQueue{
    Elements: make([]int, 0),
    Size: 0,
  }

  if n == 1 {
    return []int{0}
  }

  degrees := make(map[int]int)
  neighbors := make(map[int][]int)

  for _, edge := range edges {
    u, v := edge[0], edge[1]
    degrees[u]++
    degrees[v]++
    neighbors[u] = append(neighbors[u], v)
    neighbors[v] = append(neighbors[v], u)
  }

  for node, degree := range degrees {
    if degree == 1 {
      queue.Enqueue(node)
    }
  }

  for (n - len(visited)) > 2  {
    nodes := queue.Elements

    for _, topNode:= range nodes {
      queue.Dequeue()

      if visited[topNode] {
        continue
      }

      neighbours := neighbors[topNode]


      for _, neighbour := range neighbours {
        degrees[neighbour] -= 1
        if degrees[neighbour] ==  1 {
          queue.Enqueue(neighbour)
        }
      }

      visited[topNode] = true

    }
  }

  return queue.Elements
}

func main () {
  n := 6
 // edges := [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}
 edges := [][]int{{0,1},{0,2},{0,3},{3,4},{4,5}}

  roots := findMinHeightTrees(n, edges)

  for _, root := range roots {
    println(root)
  }
}
