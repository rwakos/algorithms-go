package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	a := []string{"pear", "amleth", "dormitory", "tinsel", "dirty room", "hamlet", "listen", "silnet"}

	fmt.Println(findAnagrams(a))
}

func findAnagrams(a []string) []string {
	used := map[int]bool{}
	r := []string{}
	for i := 0; i < len(a)-1; i++ {
		temp := []string{}
		for j := i + 1; j < len(a); j++ {
			if strings.Join(strToSlice(a[i]), "") == strings.Join(strToSlice(a[j]), "") && !used[j] {
				temp = append(temp, a[j])
				used[j] = true
			}
		}
		if len(temp) > 0 {
			temp = append([]string{a[i]}, temp...)
			val := strings.Join(temp, ",")
			r = append(r, val)
		}
	}
	return r
}

func strToSlice(s string) []string {
	s = strings.Replace(s, " ", "", -1)
	r := []string{}
	for _, v := range s {
		r = append(r, string(v))
	}
	sort.Strings(r)
	return r
}
