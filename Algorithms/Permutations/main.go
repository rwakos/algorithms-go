// Permutation Generator in Golang
package main

func main() {
	// d := permutations("ABCD")
	// fmt.Print(d)

	// d := permutations([]int{1, 2, 3, 4})
	// fmt.Print(d)

	str := "ABCD"
	permutations(str, 0, len(str))

}

// func join(ins []rune, c rune) (result []string) {
// 	for i := 0; i <= len(ins); i++ {
// 		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
// 	}
// 	return
// }

// func permutations(testStr string) []string {
// 	var n func(testStr []rune, p []string) []string

// 	n = func(testStr []rune, p []string) []string {
// 		if len(testStr) == 0 {
// 			return p
// 		}
// 		result := []string{}
// 		for _, e := range p {
// 			result = append(result, join([]rune(e), testStr[0])...)
// 		}
// 		return n(testStr[1:], result)

// 	}

// 	output := []rune(testStr)
// 	return n(output[1:], []string{string(output[0])})
// }

// 2 Way
// https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go

// func permutations(arr []int) [][]int {
// 	var helper func([]int, int)
// 	res := [][]int{}

// 	helper = func(arr []int, n int) {
// 		if n == 1 {
// 			tmp := make([]int, len(arr))
// 			copy(tmp, arr)
// 			res = append(res, tmp)
// 		} else {
// 			for i := 0; i < n; i++ {
// 				helper(arr, n-1)
// 				if n%2 == 1 {
// 					tmp := arr[i]
// 					arr[i] = arr[n-1]
// 					arr[n-1] = tmp
// 				} else {
// 					tmp := arr[0]
// 					arr[0] = arr[n-1]
// 					arr[n-1] = tmp
// 				}
// 			}
// 		}
// 	}
// 	helper(arr, len(arr))
// 	return res
// }

//3rd example

// Swap the i-th byte and j-th byte of the string
func swap(s string, i, j int) string {
	var result []byte
	for k := 0; k < len(s); k++ {
		if k == i {
			result = append(result, s[j])
		} else if k == j {
			result = append(result, s[i])
		} else {
			result = append(result, s[k])
		}
	}
	return string(result)
}

// Function to find all Permutations of a given string str[i:n]
// containing all distinct characters
func permutations(str string, i, n int) {
	// base condition
	if i == n-1 {
		print(str, " ")
		return
	}

	// process each character of the remaining string
	for j := i; j < n; j++ {
		// swap character at index i with current character
		str = swap(str, i, j)

		// recursion for string [i+1:n]
		permutations(str, i+1, n)

		// backtrack (restore the string to its original state)
		str = swap(str, i, j)
	}
}
