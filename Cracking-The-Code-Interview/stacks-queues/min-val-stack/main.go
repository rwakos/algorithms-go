package main

import (
	"sort"
)

func main() {

}

type stack []int

/*
FullStack simple struct
*/
type FullStack struct {
	min  int
	data stack
}

func (f *FullStack) pop() bool {
	s := f.data
	_, err := s.pop()

	sort.Ints(s)
	f.min = s[0]

	return err
}

func (f *FullStack) push(val int) {
	if len(f.data) == 0 || f.min > val {
		f.min = val
	}
	f.data = append([]int{val}, (*f).data...)
}

func (s *stack) pop() (int, bool) {
	if len(*s) == 0 {
		return -1, true
	}
	s1 := *s
	first := s1[0]
	*s = s1[1:]
	return first, false
}

func (s *stack) push(val int) {
	*s = append([]int{val}, *s...)
}
