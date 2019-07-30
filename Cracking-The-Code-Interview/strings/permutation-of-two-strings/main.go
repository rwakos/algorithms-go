package main

import (
	"sort"
	"strings"
)

func isPermutation(a string, b string) bool {

	if len(a) != len(b) {
		return false
	}

	if sortString(a) == sortString(b) {
		return true
	}
	return false
}

func sortString(s string) string {
	r := []string{}
	for _, v := range s {
		r = append(r, string(v))
	}
	sort.Strings(r)
	return strings.Join(r, "")
}
