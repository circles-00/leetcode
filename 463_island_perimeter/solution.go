package main


func countWaterSides(grid[][]int, Y int, X int) int {
  var count = 0
  xLength := len(grid[Y])
  yLength := len(grid)

  // Edges
  if X - 1 < 0 {
    count += 1
  }

  if X + 1 >= xLength {
    count += 1
  }

  if Y - 1 < 0 {
    count += 1
  }

  if Y + 1 >= yLength {
    count += 1
  }

  // Water
  if Y - 1 >= 0 && grid[Y - 1][X] == 0 {
    count += 1
  }

  if Y + 1 < yLength && grid[Y + 1][X] == 0 {
    count += 1
  }

  if X - 1 >= 0 && grid[Y][X - 1] == 0 {
    count += 1
  }

  if X + 1 < xLength && grid[Y][X + 1] == 0 {
    count += 1
  }

  return count
}

func islandPerimeter(grid [][]int) int {
  var count = 0
  for i:=0; i<len(grid); i++ {
    for j:=0; j<len(grid[i]); j++ {
      if grid[i][j] == 0 {
        continue
      }

      count += countWaterSides(grid, i, j)

    }
  }

  return count
}

func main () {
//  var input = [][]int{
//    {0,1,0,0},
//    {1,1,1,0},
//    {0,1,0,0},
//    {1,1,0,0}}
var input = [][]int{{1,0}}


    println(islandPerimeter(input))
}
