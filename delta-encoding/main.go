package main

import (
	"fmt"
)

func main() {
	a := []int{25626, 25757, 24367, 24267, 16, 100, 2, 7277}

	fmt.Println(deltaEncoding(a))
}

func deltaEncoding(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	breakChar := -128
	previous := nums[0]
	result := []int{previous}

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff > 127 || diff < -127 {
			result = append(result, breakChar)
		}
		result = append(result, diff)
	}
	return result
}
