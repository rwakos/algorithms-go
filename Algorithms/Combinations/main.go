package main

import (
	"fmt"
)

func main() {
	options := []string{"3", "5"}
	/*result :=[]string{"3333"}
	  total := 4*/

	result := []string{"333"}
	total := 3

	fmt.Println("Start")
	combinations(options, result, total, 0)

}

func combinations(opt []string, result []string, n int, pos int) {
	index := n - 1 - pos
	fmt.Println("index", index, "pos", pos)
	if index < 0 {
		fmt.Println(result)
		fmt.Println("End, found:", len(result))
		return
	}
	temp := []string{}
	for _, r := range result {
		for _, v := range opt {
			if v != string(r[index]) {
				temp = append(temp, swap(r, v, index))
			}

		}
	}
	result = append(result, temp...)
	combinations(opt, result, n, pos+1)

}

func swap(s string, val string, index int) string {
	if index == 0 {
		return fmt.Sprintf("%s%s", val, s[index+1:])
	} else if index == len(s)-1 {
		return fmt.Sprintf("%s%s", s[:index], val)
	} else {
		return fmt.Sprintf("%s%s%s", s[:index], val, s[index+1:])
	}
}
