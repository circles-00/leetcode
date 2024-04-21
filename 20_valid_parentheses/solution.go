package main

func isCorrectClosingBracket(first string, second string) bool {
  return (first == "(" && second == ")") || (first == "{" && second == "}") || (first == "[" && second == "]")
}

func isNested(s string) bool {
  return isCorrectClosingBracket(string(s[0]), string(s[len(s) - 1]))
}


func isValidNested(s string) bool {
  var isValid = true
  for i:=0; i<len(s) / 2; i++ {
    if !isCorrectClosingBracket(string(s[i]), string(s[len(s) - 1 - i])) {
      isValid = false
    }
  }
  return isValid
}

func isValidNotNested(s string) bool {
  var isValid = true

  for i:=0; i<len(s); i+=2 {
    if i == len(s) - 1 {
      continue
    }

    var currentChar = string(s[i])
    var nextChar = string(s[i+1])
    if !isCorrectClosingBracket(currentChar, nextChar) {
      isValid = false
    }
  }

  return isValid
}

func isValid(s string) bool {
  if(len(s) % 2 != 0) {
    return false
  }

  if isNested(s) {
    return isValidNested(s)
  }

  return isValidNotNested(s)
}

func main() {
  var input = "()[]{}"
  // var input = "{[()]}"

  println(isValid(input))
}
