package main

import (
	"fmt"
)

func main() {
	/*a := 5
	fmt.Printf("Leading or padded with space : |aa108|\n")
	//fmt.Printf("%*s\n", 20, "Hello World!")
	b := fmt.Sprintf("Leading or padded with space : |%*d|\n", a, 108)

	fmt.Printf(b)
	fmt.Printf("Leading or padded with space : |%*d|\n", a, 108)
	*/

	/*n := 6
	for k := 1; k <= n; k++ {
		p := fmt.Sprintf("|%*s|\n", n, buildPounds(k))
		fmt.Printf(p)
	}*/

	plusMinus([]int32{-4, 3, -9, 0, 4, 1})

}

func plusMinus(arr []int32) {
	b := []int32{0, 0, 0}
	n := int32(len(arr))
	for _, val := range arr {
		if val > 0 {
			b[0] = b[0] + 1
		} else if val < 0 {
			b[1] = b[1] + 1
		} else {
			b[2] = b[2] + 1
		}
	}

	for _, val := range b {
		fmt.Println(float32(val) / float32(n))
	}

}

func buildPounds(n int) string {
	result := ""
	c := "#"
	for i := 0; i < n; i++ {
		result = fmt.Sprintf("%s%s", c, result)
	}
	return result
}
