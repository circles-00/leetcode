package main

func tribonacci(n int) int {

  if n == 0 {
    return 0
  }

  nums := []int{0, 1, 1}

  for i:=3; i<=n; i++ {
    nextNumber := nums[0] + nums[1] + nums[2]
    nums[0] = nums[1]
    nums[1] = nums[2]
    nums[2] = nextNumber
  }

  return nums[2]
}

func main () {
  n := 25

  println(tribonacci(n))
}
