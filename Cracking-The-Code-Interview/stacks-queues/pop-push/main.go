package main

func main() {
}

type stack []int

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
