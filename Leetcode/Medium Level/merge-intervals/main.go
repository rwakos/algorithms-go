package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.com/problems/merge-intervals/submissions/
Runtime: 16 ms, faster than 33.73% of Go online submissions for Merge Intervals.
Memory Usage: 6.3 MB, less than 100.00% of Go online submissions for Merge Intervals.
*/
func main() {
	a := [][]int{{2, 6}, {8, 10}, {1, 3}, {15, 18}}
	b := [][]int{{1, 4}, {4, 5}}
	c := [][]int{{2, 6}, {4, 8}, {1, 3}, {15, 18}}

	print(merge(a)) //>> [[1,6],[8,10],[15,18]]
	print(merge(b)) //>> [[1,5]]
	print(merge(c)) //>> [[1,8],[15,18]]
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals[:], func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	return mergeRecursive(intervals, len(intervals))
}

func mergeRecursive(intervals [][]int, size int) [][]int {
	r := [][]int{}
	if len(intervals) <= 1 {
		return intervals
	}

	for i := 0; i < len(intervals)-1; i++ {
		lo := intervals[i][0]
		hi := intervals[i][1]

		if hi >= intervals[i+1][0] {
			if hi < intervals[i+1][1] {
				hi = intervals[i+1][1]
			}
			r = append(r, []int{lo, hi})
			i++
		} else {
			r = append(r, intervals[i])
		}

		if i == len(intervals)-2 {
			r = append(r, intervals[i+1])
		}
	}

	if size == len(r) {
		return r
	}

	return mergeRecursive(r, len(r))
}

func print(results [][]int) {
	fmt.Println(results)
}
