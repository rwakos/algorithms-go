package main

import (
	"fmt"
	"regexp"
	"strings"
)

/*
https://leetcode.com/problems/most-common-word/submissions/

Runtime: 0 ms, faster than 100.00% of Go online submissions for Most Common Word.
Memory Usage: 3.6 MB, less than 100.00% of Go online submissions for Most Common

*/
func main() {
	fmt.Println(mostCommonWord("Bob. hIt, baLl", []string{"bob", "hit"})) // "ball"
}

func mostCommonWord(paragraph string, banned []string) string {
	counter := map[string]int{}
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
	}
	processedParagraph := reg.ReplaceAllString(paragraph, " ")
	words := strings.Split(strings.ToLower(processedParagraph), " ")

	for _, v := range words {
		if v != " " && v != "" {
			counter[v]++
		}
	}

	r := ""
	max := 0
	found := false
	for k, v := range counter {
		if !inArray(banned, k) {
			if !found {
				r = k
				max = counter[k]
				found = true
			} else if v > max {
				r = k
				max = counter[k]
			}
		}
	}
	return r

}

func inArray(a []string, word string) bool {
	for _, v := range a {
		if word == v {
			return true
		}
	}
	return false
}
