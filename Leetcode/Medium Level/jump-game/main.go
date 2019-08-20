package main

import (
	"fmt"
)

func main() {
	//fmt.Println(canJump([]int{0}))             //True
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // True
	//fmt.Println(canJump([]int{1, 0, 2}))       // False
	//fmt.Println(canJump([]int{2, 0}))          // True

}

func canJump(nums []int) bool {
	max := len(nums)

	if max <= 1 {
		return true
	}

	pos := 0
	for i := 0; i < max; i++ {
		jump := nums[pos]
		pos += jump - 1

		fmt.Println(jump, pos)

		if (pos + 1) == max {
			return true
		} else if (pos + 1) > max {
			return false
		}
	}
	return false
}
