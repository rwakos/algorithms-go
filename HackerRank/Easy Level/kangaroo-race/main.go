package main

import "fmt"

func main() {
	fmt.Println(kangaroo(5, 2, 6, 1))
}

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	/*
			Hack way
			if ((x1 - x2) % (v1 - v2)) == 0 {
		        return "YES"
		    }*/
	if v1 <= v2 {
		return "NO"
	}

	for x1 < x2 {
		if x1+v1 == x2+v2 {
			return "YES"
		}
		x1 += v1
		x2 += v2
	}

	return "NO"
}
