package main

import (
	"fmt"
)

/*
https://leetcode.com/problems/spiral-matrix/submissions/
Runtime: 0 ms, faster than 100.00% of Go online submissions for Spiral Matrix.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Spiral Matrix.
*/
func main() {

	m1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	fmt.Println(spiralOrder(m1))
	fmt.Println("expected:", "\n[1,2,3,6,9,8,7,4,5]")

	m2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(m2))
	fmt.Println("expected:", "\n[1 2 3 4 8 12 16 15 14 13 9 5 6 7 11 10]")
}

func spiralOrder(matrix [][]int) []int {
	var r []int
	rows := len(matrix)
	if rows == 0 {
		return []int{}
	}
	cols := len(matrix[0])
	end := rows * cols

	direction := "top_to_right"
	row := 0
	col := 0
	for i := 0; i < end; i++ {
		if direction == "top_to_right" {
			for j := 0; j < cols; j++ {
				r = append(r, matrix[row][col])
				i++
				col++
			}
			col--
			i--
			direction = "top_to_bottom"
		} else if direction == "top_to_bottom" {
			for j := 0; j < rows-1; j++ {
				row++
				r = append(r, matrix[row][col])
				i++
			}
			direction = "bottom_to_left"
			i--
		} else if direction == "bottom_to_left" {
			for j := 0; j < cols-1; j++ {
				col--
				r = append(r, matrix[row][col])
				i++
			}
			direction = "bottom_to_top"
			i--
		} else if direction == "bottom_to_top" {
			for j := 0; j < rows-2; j++ {
				row--
				r = append(r, matrix[row][col])
				i++
			}
			direction = "top_to_right"

			i--
			rows -= 2
			cols -= 2
			col++
		}
	}

	return r
}
