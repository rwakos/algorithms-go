package main

import (
	"fmt"
	"sort"
)

/*
	Find the duplicates between arrays. O(n)
	Both arrays are sorted and have the different length
	Note: use the same approach, jus watch the overflow of both arrays
*/
func main() {
	a := []int{13, 27, 35, 40, 49, 55, 58, 72, 83}
	b := []int{17, 35, 39, 40, 55, 58, 72}
	fmt.Println(findDuplicates(a, b))

	a = []int{17, 35, 39, 40, 55, 58, 72}
	b = []int{13, 27, 35, 40, 49, 55, 58, 72, 83}
	fmt.Println(findDuplicates(a, b))
}

func findDuplicates(a []int, b []int) []int {
	map1 := map[int]int{}
	map2 := map[int]int{}
	result := map[int]int{}

	top := len(b)

	if len(a) > len(b) {
		top = len(a)
	}

	for i := 0; i < top; i++ {
		if i < len(a) && i+1 < len(b) && a[i] == b[i] {
			result[a[i]] = a[i]
		} else if i < len(a) && a[i] == map2[a[i]] {
			result[a[i]] = a[i]
		} else if i < len(b) && b[i] == map1[b[i]] {
			result[b[i]] = b[i]
		}
		if i < len(a) {
			map1[a[i]] = a[i]
		}
		if i < len(b) {
			map2[b[i]] = b[i]
		}
	}

	r := []int{}
	for k := range result {
		r = append(r, k)
	}
	sort.Ints(r) //Just to test if we get the same result
	return r
}
