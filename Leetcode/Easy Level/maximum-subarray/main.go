package main

import "fmt"

/*
Runtime: 88 ms, faster than 7.95% of Go online submissions for Maximum Subarray.
Memory Usage: 3.3 MB, less than 57.14% of Go online submissions for Maximum
*/
func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a)) //returns 6

	a = []int{-1}
	fmt.Println(maxSubArray(a)) //returns -1

	a = []int{-1, -2}
	fmt.Println(maxSubArray(a)) //returns -1
}

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := 0
	max := -2147483648
	for i := 0; i < len(nums); i++ {
		sum = nums[i]
		if sum > max {
			max = sum
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}
	return max
}
