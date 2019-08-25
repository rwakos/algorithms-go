package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{0}))             //True
	fmt.Println(canJump([]int{2, 3, 1, 1, 4})) // True
	fmt.Println(canJump([]int{1, 0, 2}))       // False
	fmt.Println(canJump([]int{2, 0}))          // True
	fmt.Println(canJump([]int{2, 0, 0}))       // True
	fmt.Println(canJump([]int{2, 5, 0, 0}))    // True

}

func canJump(nums []int) bool {
	max := len(nums)

	if max <= 1 {
		return true
	}
	pos := nums[0]

	for i := 0; i < max; i++ {
		if pos >= len(nums)-1 {
			return true
		}

		jump := nums[pos]

		if pos+jump >= len(nums)-1 {
			return true
		}
		if nums[pos+jump] != 0 {
			pos += jump
		}
	}
	return false
}
