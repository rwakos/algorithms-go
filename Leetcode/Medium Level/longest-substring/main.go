package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbb"))   // 1 >> b
	fmt.Println(lengthOfLongestSubstring("pwwkew")) // 3 >> wke
}

func lengthOfLongestSubstring(s string) int {
	max := 0
	sum := 0
	for i := range s {
		m := make(map[string]int)
		sum = 0
		for _, v1 := range s[i:] {
			if m[string(v1)] == 0 {
				m[string(v1)] = 1
				sum++
			} else {
				if sum > max {
					max = sum
				}
				break
			}
			if sum != 0 && sum > max {
				max = sum
			}
		}
	}

	return max
}
