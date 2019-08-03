package main

import (
	"fmt"
	"strings"
)

func main() {
	m1 := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	printMatrix(m1)
	printMatrix(rotateMatrix(m1))
}

func rotateMatrix(m [][]string) [][]string {
	newMatrix := make([][]string, len(m))
	for i := 0; i < len(m); i++ {
		newMatrix[i] = make([]string, len(m))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			newMatrix[j][len(m)-i-1] = m[i][j]
		}
	}

	return newMatrix
}

func printMatrix(p [][]string) {
	for i := 0; i < len(p); i++ {
		fmt.Println(p[i])
	}
	fmt.Println("-------------------------")

}

func compareMatrix(m1 [][]string, m2 [][]string) bool {
	for i := 0; i < len(m1); i++ {
		if strings.Join(m1[i], "") != strings.Join(m2[i], "") {
			return false
		}
	}
	return true
}
