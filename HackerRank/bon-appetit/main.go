package main

/*
https://www.hackerrank.com/challenges/bon-appetit/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
*/
import "fmt"

func main() {

}

func bonAppetit(bill []int32, k int32, b int32) string {
	sum := int32(0)
	for _, v := range bill {
		sum += int32(v)
	}

	fullBill := (sum / int32(2)) - b - (bill[k] / int32(2))

	if fullBill == 0 {
		return "Bon Appetit"
	}

	return fmt.Sprintf("%d", (-fullBill))
}
