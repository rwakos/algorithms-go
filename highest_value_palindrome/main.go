package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(highestValuePalindrome("092282", 6, 3))
}

func highestValuePalindrome(s string, n int32, k int32) string {
	/*var result = make([]string, n)
	var i int
	//typeNumber := "even"
	max := int(n / 2)
	var change int32

	if n%2 > 0 {
		result[max+1] = string(s[max+1])
		//typeNumber = "odd"
	}

	fmt.Println(s)
	diff := k - change
	if diff > 1 {
		for diff-1 > 0 {
			if int32(s[i]) != int32(s[len(s)-i-1]) {
				result[i] = "9"
				result[len(s)-i-1] = "9"
				if string(s[i]) != "9" {
					diff--
				}
				if string(s[len(s)-i-1]) != "9" {
					diff--
				}
			} else {
				result[i] = string(s[i])
				result[len(s)-i-1] = string(s[len(s)-i-1])
			}
			i++
		}
		s = strings.Join(result, "")
	}

	for i := 0; i < max; i++ {
		if int32(s[i]) < int32(s[len(s)-i-1]) {
			result[i] = string(s[len(s)-i-1])
			result[len(s)-i-1] = string(s[len(s)-i-1])
			diff--
		} else if int32(s[i]) > int32(s[len(s)-i-1]) {
			result[i] = string(s[i])
			result[len(s)-i-1] = string(s[i])
			diff--
		} else {
			result[i] = string(s[i])
			result[len(s)-i-1] = string(s[i])
		}
	}

	if int32(diff) == 0 {
		return strings.Join(result, "")
	}

	fmt.Println(result)
	return "-1"

	/*if (change == 0 && n%2 == 0 && k%2 > 0) || int32(change) > k {
		return "-1"
	}
	return strings.Join(result, "")*/
	var result = make([]string, n)
	var max = int(n)
	var moves = int(k)
	var lo int
	hi := max - 1
	var diff int

	for i := 0; i < max/2; i++ {
		j := max - i - 1
		if s[i] != s[j] {
			diff++
		}
		result[i] = string(s[i])
		result[j] = string(s[j])
	}

	if diff > moves {
		return "-1"
	}

	for hi >= lo {
		if k <= 0 {
			break
		}

		if s[lo] == s[hi] {
			if moves > 1 && (moves-2) >= diff && string(s[lo]) != "9" {
				result[lo] = "9"
				result[hi] = "9"
				moves -= 2
			}
		} else {
			if moves > 1 && (moves-2) >= diff-1 {
				if string(s[lo]) != "9" {
					result[lo] = "9"
					moves--
				}
				if string(s[hi]) != "9" {
					result[hi] = "9"
					moves--
				}
			} else {
				if s[lo] > s[hi] {
					result[hi] = string(result[lo])
				} else {
					result[lo] = string(s[hi])
				}
				moves--
			}
			diff--
		}
		if lo == hi && moves > 0 {
			result[lo] = "9"
			moves--
		}
		lo++
		hi--
	}

	if isPalindrome(strings.Join(result, "")) {
		return strings.Join(result, "")
	}

	return "-1"
}

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			fmt.Println(s[i], s[len(s)-i-1])
			return false
		}
	}
	return true
}
