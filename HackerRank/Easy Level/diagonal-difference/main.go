package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}

func diagonalDifference(arr [][]int32) int32 {
	max := int(len(arr))
	sumLeft := int32(0)
	sumRight := int32(0)
	for i := 0; i < max; i++ {
		sumLeft += arr[i][i]
		sumRight += arr[i][max-i-1]
	}

	return int32(math.Abs(float64(sumLeft) - float64(sumRight)))
}
