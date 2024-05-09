package main

func squareIsWhite(coordinates string) bool {
	lettersColors := map[byte]string{
		97:  "black",
		98:  "white",
		99:  "black",
		100: "white",
		101: "black",
		102: "white",
		103: "black",
		104: "white",
	}

	if lettersColors[coordinates[0]] == "black" {
		return coordinates[1]%2 == 0
	}

	if lettersColors[coordinates[0]] == "white" {
		return coordinates[1]%2 != 0
	}

	return false
}

func main() {
	coordinates := "h3"

	println(squareIsWhite(coordinates))
}
