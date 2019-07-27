package main

import (
	"fmt"
)

func main() {
	a := []int{5, 6, 7, 4, 3}
	fmt.Println(twoSum(a, 9))
}

func twoSum(nums []int, target int) []int {
	comps := map[int]int{}

	for i, val := range nums {
		comp := target - val

		index, ok := comps[comp]
		if ok {
			return []int{index, i}
		}

		comps[val] = i
	}

	return []int{}
}
