package main

import (
	"fmt"
	"strings"
)

func main() {
	decentNumber(1)  // -1
	decentNumber(3)  // 555
	decentNumber(5)  // 33333
	decentNumber(11) // 55555533333
}

func decentNumber(n int32) {
	if n <= 2 {
		fmt.Println(-1)
		return
	}

	z := n
	for {
		if z%3 != 0 {
			z -= 5
			continue
		}

		if z < 0 {
			fmt.Println(-1)
			break
		} else {
			printResults(z, n-z)
			break
		}
	}
}

func printResults(five int32, three int32) {
	r := ""

	if five > 0 {
		r = strings.Repeat("5", int(five))
	}

	if three > 0 {
		r = fmt.Sprintf("%s%s", r, strings.Repeat("3", int(three)))
	}

	fmt.Println(r)
}
