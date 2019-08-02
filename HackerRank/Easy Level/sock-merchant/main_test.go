package main

import (
	"fmt"
	"testing"
)

func TestSockMerchant(t *testing.T) {
	a := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	response := sockMerchant(int32(len(a)), a)

	if response != 3 {
		t.Errorf("FAIL, should return 3, returned: %d", response)
	} else {
		fmt.Println("PASS - Case 1")
	}

	a = []int32{1, 2, 1, 2, 1, 3, 2}
	response = sockMerchant(int32(len(a)), a)

	if response != 2 {
		t.Errorf("FAIL, should return 2, returned: %d", response)
	} else {
		fmt.Println("PASS - Case 2")
	}
	fmt.Println("---------------------------------")
}
