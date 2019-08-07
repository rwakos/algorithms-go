package main

import "fmt"

/*
https://www.hackerrank.com/challenges/utopian-tree/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
*/

func main() {
	fmt.Println(utopianTree(4)) // Should be 7
}

func utopianTree(n int32) int32 {
	acum := int32(0)
	for i := int32(0); i <= n; i++ {
		if i == 0 {
			acum = 1
		} else if i%2 > 0 {
			acum = acum * 2
		} else {
			acum++
		}
	}
	return acum
}
