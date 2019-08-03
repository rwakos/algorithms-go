package main

import "fmt"

func main() {

}

func compressString(s string) string {
	if len(s) == 0 {
		return ""
	}
	compressedString := ""
	currentChar := string(s[0])
	compressedString = fmt.Sprintf("%s%s", compressedString, currentChar)
	counter := 1

	for i := 1; i < len(s); i++ {
		if currentChar != string(s[i]) {
			if counter > 1 {
				compressedString = fmt.Sprintf("%s%d", compressedString, counter)
			}

			counter = 1
			currentChar = string(s[i])

			compressedString = fmt.Sprintf("%s%s", compressedString, currentChar)

		} else {
			counter++
		}
	}

	if counter > 1 {
		compressedString = fmt.Sprintf("%s%d", compressedString, counter)
	}

	return compressedString
}
