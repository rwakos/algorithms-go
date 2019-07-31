package main

import (
	"fmt"
	"math/big"
)

func main() {
	fmt.Println(smallestString64(64))
}

func smallestString64(weight int64) string {
	response := ""
	m, c := getDictionary64(25)
	max := int(([]rune(c))[0])
	acum := weight
	for i := max; i > 64; i-- {
		key := string(i)
		v := m[key]
		added := int(acum / v)

		if added > 0 {
			acum = acum % v
			for j := 0; j < added; j++ {
				response = fmt.Sprintf("%s%s", key, response)
			}
		}

		if acum == 0 {
			return response
		}
	}

	return response
}

func getDictionary64(w int64) (map[string]int64, string) {
	m := map[string]int64{}
	m["A"] = int64(1)
	index := 1
	prev := "A"
	for i := 1; i < 26; i++ {
		index = int(([]rune("A"))[0]) + i
		prev = string(index - 1)
		leftside := (int64(index)-65+1)*m[prev] + m[prev]
		if leftside > w {
			return m, prev
		}
		m[string(index)] = leftside
	}
	return m, string(index)
}

func getDictionary(w int64) (map[string]*big.Int, string) {
	m := map[string]*big.Int{}
	wBig := big.NewInt(w)
	m["A"] = big.NewInt(1)
	index := 1
	prev := "A"
	for i := 1; i < 26; i++ {
		index = int(([]rune("A"))[0]) + i
		prev = string(index - 1)

		leftside := big.NewInt(int64(index) - 65 + 1)
		leftside.Mul(leftside, m[prev])
		leftside.Add(leftside, m[prev])

		if leftside.Cmp(wBig) == 1 {
			return m, prev
		}

		m[string(index)] = leftside
	}

	return m, string(index)
}
