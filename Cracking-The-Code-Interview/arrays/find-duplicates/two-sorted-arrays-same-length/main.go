package main

import "fmt"

/*
	Find the duplicates between arrays. O(n)
	Both arrays are sorted and have the same length
*/
func main() {
	a := []int{13, 27, 35, 40, 49, 55, 60}
	b := []int{17, 35, 39, 40, 55, 58, 60}
	fmt.Println(findDuplicates(a, b))
}

func findDuplicates(a []int, b []int) []int {
	map1 := map[int]int{}
	map2 := map[int]int{}
	result := map[int]int{}

	for i := 0; i < len(a); i++ {
		if a[i] == b[i] {
			result[a[i]] = a[i]
		} else if a[i] == map2[a[i]] {
			result[a[i]] = a[i]
		} else if b[i] == map1[b[i]] {
			result[b[i]] = b[i]
		}
		map1[a[i]] = a[i]
		map2[b[i]] = b[i]
	}

	r := []int{}
	for k := range result {
		r = append(r, k)
	}
	return r
}
