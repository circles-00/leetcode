package main

func findSmallestString(strs []string) string {
  var smallest=strs[0]

  for _, string := range strs {
    if len(string) < len(smallest){
      smallest = string
    }
  }

  return smallest
}

func longestCommonPrefix(strs []string) string {
  smallestString := findSmallestString(strs)

  var commonPrefix = ""
  for i, char := range smallestString {
    var addChar = true
    for j:=0; j<len(strs); j++ {
      if rune(strs[j][i]) != char {
        addChar = false
      }
    }

    if !addChar {
      break
    }

    commonPrefix = commonPrefix + string(char)
  }

  return commonPrefix
}

func main () {
  // var strs = []string{"cir","car"}
  var strs = []string{"flower","flow","flight"}
  println(longestCommonPrefix(strs))
}

