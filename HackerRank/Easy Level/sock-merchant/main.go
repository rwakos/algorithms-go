package main

func main() {

}

/*
https://www.hackerrank.com/challenges/sock-merchant/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
O(n)
*/
func sockMerchant(n int32, ar []int32) int32 {
	m := map[int32]int{}
	count := int32(0)
	for _, v := range ar {
		m[v]++

		if m[v] == 2 {
			count++
			m[v] = 0
		}
	}
	return count
}
