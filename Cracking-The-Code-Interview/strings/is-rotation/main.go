package main

import "fmt"

func main() {
	if !isRotation("water", "terwa") {
		fmt.Println("error")
	}
}

func isRotation(original string, possible string) bool {
	if len(original) == 0 {
		return false
	}

	if len(original) != len(possible) {
		return false
	}

	positionsOriginal := findIndexesOfFirstChar(string(original[0]), original)
	positions := findIndexesOfFirstChar(string(original[0]), possible)

	if len(positionsOriginal) != len(positions) {
		return false
	}

	flag := false
	i := 0
	j := 0

	for _, v := range positions {
		flag = true
		j = 0
		i = v
		for k := 0; k < len(original); k++ {
			if original[j] != possible[i] {
				flag = false
				break
			}

			if i == len(original)-1 {
				i = -1
			}
			i++
			j++
		}

		if flag {
			return true
		}
	}

	return false
}

func findIndexesOfFirstChar(firstChar string, o string) []int {
	r := []int{}

	for i := 0; i < len(o); i++ {
		if string(o[i]) == firstChar {
			r = append(r, i)
		}
	}
	return r
}
