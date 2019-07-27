package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "baab"
	fmt.Println(superReducedString(s))
}

func superReducedString(s string) string {
	temp := s

	if len(s) == 1 {
		return s
	}
	if len(s) == 0 {
		return "Empty String"
	}

	for _, v := range s {
		doubleChr := fmt.Sprintf("%s%s", string(v), string(v))
		temp = strings.Replace(temp, doubleChr, "", -1)
	}

	if len(s) > len(temp) {
		return superReducedString(temp)
	}

	return temp
}
