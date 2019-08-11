package main

import (
	"fmt"
	"strconv"
)

/*
100% Faster
https://leetcode.com/problems/reverse-integer/
*/
func main() {

}

func reverse(x int) int {
	var strInt string
	var sign string
	var response string
	strInt = fmt.Sprintf("%d", x)

	if x < 0 {
		sign = "-"
		strInt = fmt.Sprintf("%d", -x)
	}

	for j := len(strInt) - 1; j >= 0; j-- {
		response = fmt.Sprintf("%s%s", response, string(strInt[j]))
	}

	r, _ := strconv.Atoi(fmt.Sprintf("%s%s", sign, response))

	if r > 2147483647 || r < -2147483647 {
		return 0
	}
	return r
}
