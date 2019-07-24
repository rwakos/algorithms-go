package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(minimumNumber(3, "Ab1"))
}

func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	numbers := "0123456789"
	lowerCase := "abcdefghijklmnopqrstuvwxyz"
	upperCase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters := "!@#$%^&*()-+"

	score := map[string]bool{}

	for _, v := range password {
		if strings.Contains(numbers, string(v)) {
			score["numbers"] = true
		}

		if strings.Contains(lowerCase, string(v)) {
			score["lowerCase"] = true
		}

		if strings.Contains(upperCase, string(v)) {
			score["upperCase"] = true
		}

		if strings.Contains(specialCharacters, string(v)) {
			score["specialCharacters"] = true
		}
	}

	count := int32(0)
	for _, v := range score {
		if v {
			count++
		}
	}
	count = 4 - count

	if n >= 6 {
		return count
	}
	return 6 - n

}
