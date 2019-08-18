package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(lengthOfLastWord("Hello World")) // 5
}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Length of Last Word.
Memory Usage: 2.2 MB, less than 33.33% of Go online submissions for Length of Last Word.
*/

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	w := strings.Split(strings.Trim(s, " "), " ")

	return len(w[len(w)-1])
}
