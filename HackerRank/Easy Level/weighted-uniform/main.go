package main

/*
https://www.hackerrank.com/challenges/weighted-uniform-string/problem
*/
import "fmt"

func main() {
	fmt.Println(weightedUniformStrings("aaabbbbcccddd", []int32{9, 7, 8, 12, 5}))
}

func weightedUniformStrings(s string, queries []int32) []string {
	r := []string{}
	values := map[int32]bool{}
	prev := int32(0)
	count := int32(0)

	for _, v := range s {
		value := int32(v) - 96
		if prev == value {
			count += value
			values[count] = true
		} else {
			count = value
			values[value] = true
			prev = value
		}
	}

	fmt.Println(values)
	for _, v := range queries {
		_, ok := values[v]
		if ok {
			r = append(r, "Yes")
		} else {
			r = append(r, "No")
		}
	}
	return r
}
