package main

import (
	"os"
)

func main() {
}

type stack []int

func (s *stack) sort() {
	var r stack

	for i := 0; i < len(*s); i++ {
		v := (*s)[i]
		newIndex := minMax(&r, v)
		if !r.pushAtIndex(newIndex, v) {
			os.Exit(2)
		}
	}

	*s = r
}

func (s *stack) pushAtIndex(index int, value int) bool {
	var temp stack

	if index > len(*s) {
		return false
	}

	if index == -1 {
		s.push(value)
		return true
	}

	for i := 0; i < index; i++ {
		popValue, err := s.pop()
		if !err {
			temp = append([]int{popValue}, temp...)
		}
	}

	s.push(value)

	for _, v := range temp {
		s.push(v)
	}
	return true
}

func (s *stack) popAtIndex(index int) (int, bool) {
	var temp stack

	if index > len(*s)-1 {
		return -1, true
	}

	for i := 0; i < index; i++ {
		popValue, err := s.pop()
		if !err {
			temp = append([]int{popValue}, temp...)
		}
	}

	current, _ := s.pop()

	for _, v := range temp {
		s.push(v)
	}
	return current, false
}

func minMax(s *stack, value int) int {
	if len(*s) == 0 {
		return -1
	}
	min := -1
	index := len(*s)
	found := false
	for i, v := range *s {
		if v > value {
			if !found {
				min = v
				index = i
				found = true
			} else if v < min {
				min = v
				index = i
			}
		}
	}
	return index
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
