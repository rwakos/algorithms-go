package main

import "fmt"

/*
	Create an ATM machine, that allows you to withdraw money if available and print a receipt
	of the bills and amount, if not print out Error
	Initial load 3 bills of 50, 2 of 20, 1 of 10, and 2 of 5
	Note: you have 30 min to do this...
*/
func main() {
	//Initial load
	balance := map[int]int{50: 3, 20: 2, 10: 1, 5: 2}

	//withdraw 75
	balance = withdrawMoney(balance, 75)

	//withdraw 3, should message error
	balance = withdrawMoney(balance, 3)

}

func withdrawMoney(balance map[int]int, amount int) map[int]int {
	withdraw := map[int]int{}
	availableBill := []int{50, 20, 10, 5}

	for _, bill := range availableBill {
		qtyAvailable := balance[bill]

		if amount > 0 {
			substract := int(amount / bill)

			if substract > 0 {
				if substract > qtyAvailable {
					amount = amount - (qtyAvailable * bill)
					withdraw[bill] = qtyAvailable
				} else {
					amount = amount - (substract * bill)
					withdraw[bill] = substract
				}
			}
		}

		if amount == 0 {
			break
		}
	}

	if amount != 0 {
		fmt.Println("error")
		return balance
	}

	printReceipt(withdraw)
	return makeWithdraw(balance, withdraw)
}

func printReceipt(withdraw map[int]int) {
	availableBill := []int{50, 20, 10, 5}
	fmt.Println("--------------------")
	fmt.Println("------RECEIPT-------")
	fmt.Println("--------------------")
	for _, bill := range availableBill {
		if withdraw[bill] > 0 {
			fmt.Println(fmt.Sprintf("%d bills of %d", withdraw[bill], bill))
		}
	}
	fmt.Println("--------------------")
}

func makeWithdraw(balance map[int]int, withdraw map[int]int) map[int]int {
	for k, qty := range withdraw {
		balance[k] = balance[k] - qty
	}
	return balance
}
