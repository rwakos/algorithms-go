package main

import "fmt"

func main() {
	fmt.Println(camelcase("howAreYouStudiyin"))
	//Should print 4
}

func camelcase(s string) int32 {
	count := 1
	for _, v := range s {
		c := v
		if c < 97 {
			count++
		}
	}
	return int32(count)
}
