package main

import "fmt"

/*
	Find the duplicates between arrays. O(n)
	Both arrays are sorted and have different length, and you can not use hash maps
*/
func main() {
	a := []int{13, 27, 35, 40, 49, 55, 60, 62, 64, 70}
	b := []int{17, 35, 39, 40, 55, 58, 60, 70}
	findDuplicates(a, b)
}

func findDuplicates(a []int, b []int) {
	i := 0
	j := 0

	for k := 0; k < len(a)+len(b); k++ {
		eval := false
		if i < len(a) && j < len(b) {
			eval = true
		}
		if eval && a[i] == b[j] {
			fmt.Print(a[i], " ")
			i++
			j++
		} else if eval && a[i] < b[j] {
			i++
		} else if eval && b[j] < a[i] {
			j++
		}
	}
}
