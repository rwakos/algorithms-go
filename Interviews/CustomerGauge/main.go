package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := makeTimestamp()
	findPrimeNumbersWhithin(100)
	ended := makeTimestamp()
	v1 := ended - start
	fmt.Println("Time V1", v1)

	start = makeTimestamp()
	findPrimeNumbersWhithinV2(100)
	ended = makeTimestamp()
	v2 := ended - start
	fmt.Println("Time V2", v2)

	fmt.Printf("V2 is %f faster", float64(v1)/float64(v2))
}

func findPrimeNumbersWhithin(n int) {
	if n < 2 {
		fmt.Println("No primes")
	}

	for i := 2; i <= n; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println("")
}

func findPrimeNumbersWhithinV2(n int) {
	if n < 2 {
		fmt.Println("No primes")
	}

	sqrt := int(math.Sqrt(float64(n)))

	for i := 2; i <= n; i++ {
		flag := true
		for j := 2; j <= sqrt; j++ {
			if i%j == 0 && i != j {
				flag = false
				break
			}
		}

		if flag {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Println("")
}

func makeTimestamp() int64 {
	return time.Now().UnixNano()
}
