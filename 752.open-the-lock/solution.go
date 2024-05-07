package main


type Queue interface {
  Enqueue(item string)
  Dequeue()
  Peek() string
}

type SliceQueue struct {
  Elements []string
  Size int
}

func (q * SliceQueue) Enqueue(item string) {
  q.Elements = append(q.Elements, item)
  q.Size += 1
}

func (q * SliceQueue) Dequeue() {
  q.Elements = q.Elements[1:]
  q.Size -= 1
}

// Returns empty string if the queue is empty
func (q SliceQueue) Peek() string {
  if len(q.Elements) > 0 {
    return q.Elements[0]
  }

  return ""
}

func findNeighbouringNodes(currentNode string) []string {
  neighbours := make([]string, 0)
  for index, char := range currentNode {
    digit := int(char)
    runeNodeForward := []rune(currentNode)
    runeNodeForward[index] = rune(digit + 1)

    runeNodeBackward := []rune(currentNode)
    runeNodeBackward[index] = rune(digit - 1)


    // Figured out digit is casted to an ASCII value, so use ASCII table for this one
    if digit == 48 {
      runeNode := []rune(currentNode)
      runeNode[index] = rune(57)

      neighbours = append(neighbours, string(runeNode))
    }

    if digit == 57 {
      runeNode := []rune(currentNode)
      runeNode[index] = rune(48)

      neighbours = append(neighbours, string(runeNode))
    }

    if digit >= 48 && digit <= 57 {
      if digit != 57 {
        neighbours = append(neighbours, string(runeNodeForward))
      }
      neighbours = append(neighbours, string(runeNodeBackward))
    }
  }

  return neighbours
}

func openLock(deadends []string, target string) int {
  startingNode := "0000"

  visited := make(map[string]bool)
  distances := make(map[string]int)
  queue := SliceQueue{
    Elements: make([]string, 0),
    Size: 0,
  }

  // We're gonna have to check for deadends a lot, so map is way better here, so we can check in constant time
  deadendsMap := make(map[string]bool)

  for _, deadend := range deadends {
    deadendsMap[deadend] = true
  }

  visited[startingNode] = true
  queue.Enqueue(startingNode)

  if deadendsMap[startingNode] {
    return -1
  }

  if startingNode == target {
    return 0
  }

  for queue.Size > 0 {
    topNode := queue.Peek()

    if topNode == target {
      break
    }

    neighbours := findNeighbouringNodes(topNode)
    for _, neighbour := range neighbours {
      if !deadendsMap[neighbour] && !visited[neighbour] {
        visited[neighbour] = true
        queue.Enqueue(neighbour)
        distances[neighbour] = distances[topNode] + 1
      }
    }

    queue.Dequeue()
  }

  if distances[target] == 0 {
    return -1
  }

  return distances[target]
}

func main () {
 // deadends := []string{
 //   "0201","0101","0102","1212","2002",
 // }

 // target := "0202"
 deadends := []string{
   "8887","8889","8878","8898","8788","8988","7888","9888",
 }
 target := "8888"

  println(openLock(deadends, target))
   neighbours := findNeighbouringNodes("8888")
   for _, neighbour := range neighbours {
     println(neighbour)
   }
}
