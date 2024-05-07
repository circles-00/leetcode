package main

func absoluteDiff (x, y int) int {
  if y > x {
    return y - x
  }

  return x - y
}

func longestIdealString(s string, k int) int {
}

func main () {
  // s := "acfgbd"
  s := "lkpkxcigcs"
  k := 6

  println(longestIdealString(s, k))
}
