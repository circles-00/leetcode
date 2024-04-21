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

func validPath(n int, edges[][]int, source int, destination int) bool {
  visited := make(map[int]bool)

  queue := SliceQueue{
    Elements: make([]int, 0),
    Size: 0,
  }

  visited[source] = true
  queue.Enqueue(source)

  for queue.Size > 0 {
    topNode := queue.Peek()
    if topNode == destination {
      return true
    }

    for i:=0; i < len(edges); i++ {
      if edges[i][0] == topNode && !visited[edges[i][1]] {
        queue.Enqueue(edges[i][1])
        visited[edges[i][1]] = true
      }

      if edges[i][1] == topNode && !visited[edges[i][0]] {
        queue.Enqueue(edges[i][0])
        visited[edges[i][0]] = true
      }
    }

    queue.Dequeue()
  }

  return visited[destination]
}

func main() {
  n := 3
  edges := [][]int{
    {0, 1},
    {1, 2},
    {2, 0},
  }
  source := 0
  destination := 2

  println(validPath(n, edges, source, destination))
  //neighbours := findNeighbouringNodes(edges, 0)

  //  for i:=0; i<len(neighbours); i++ {
  //    println(neighbours[i])
  //  }
}
