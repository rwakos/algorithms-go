package main

import (
	"fmt"
	"strings"
)

func main() {
	myMap := buildMap("programmer")
	fmt.Println(programmerStrings2("progxrammerrxproxgrammer", myMap))
	fmt.Println(programmerStrings2("programmerprogrammer", myMap))
}

func programmerStrings(s string) int32 {
	// Write your code here
	allowed := initMe()

	if len(strings.Replace(s, "programmer", "", -1)) == 0 {
		return int32(0)
	}

	pos := 0
	count := int32(0)
	finished := false
	for _, v := range s {
		if string(v) == allowed[pos] {
			pos++
			if pos >= len(allowed) {
				pos = 0
				finished = true
			}
		}
		if finished && pos == 0 {
			count++
		}
	}
	return count
}

func initMe() []string {
	return []string{"p", "r", "o", "g", "r", "a", "m", "m", "e", "r"}
}

func programmerStrings2(s string, mapped map[string]int) int32 {
	mappedCopy := mapped
	startCount := false
	counter := int32(0)

	for _, v := range s {
		qty, found := mapped[string(v)]
		if found {
			mapped[string(v)] = qty - 1
		} else if startCount {
			counter++
		}

		if isFinished(mapped) {
			mapped = mappedCopy
			startCount = true
		}
	}

	return counter
}

func isFinished(acum map[string]int) bool {
	for _, v := range acum {
		if v > 0 {
			return false
		}
	}
	return true
}

func buildMap(word string) map[string]int {
	acum := map[string]int{}

	for _, v := range word {
		acum[string(v)] = acum[string(v)] + 1
	}
	return acum
}
