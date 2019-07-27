package main

import (
	"fmt"
)

func main() {
	a := []int32{5, 6, 7}
	b := []int32{3, 6, 10}
	fmt.Println(compareTriplets(a, b))
}

func compareTriplets(a []int32, b []int32) []int32 {
	r := []int32{0, 0}

	for i, v := range a {
		if v > b[i] {
			r[0]++
		} else if v < b[i] {
			r[1]++
		}
	}
	return r
}
