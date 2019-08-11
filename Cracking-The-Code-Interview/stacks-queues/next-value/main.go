package main

func main() {
}

type stack []int

/*
Find the index where the value should be placed
*/
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

func (s *stack) push(val int) {
	*s = append([]int{val}, *s...)
}
