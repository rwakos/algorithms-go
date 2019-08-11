package main

import "fmt"

/*
65% faster
https://leetcode.com/problems/palindrome-number/
*/

func main() {
	fmt.Println(isPalindrome(121))  //True
	fmt.Println(isPalindrome(-121)) //False
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	atoi := fmt.Sprintf("%d", x)
	n := int(len(atoi) / 2)
	for i := 0; i < n; i++ {
		if atoi[i] != atoi[len(atoi)-1-i] {
			return false
		}
	}
	return true
}
