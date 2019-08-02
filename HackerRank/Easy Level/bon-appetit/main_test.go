package main

import (
	"fmt"
	"testing"
)

func TestBonAppetit(t *testing.T) {
	response := bonAppetit([]int32{3, 10, 2, 9}, 1, 12)

	//Case 1
	if response != "5" {
		t.Errorf("FAILED: this should be 5, %s given", response)
	} else {
		fmt.Println("PASS - Case 1")
	}

	//Case 2
	response = bonAppetit([]int32{3, 10, 2, 9}, 1, 7)

	if response != "Bon Appetit" {
		t.Errorf("FAILED: this should be 5, %s given", response)
	} else {
		fmt.Println("PASS - Case 2")
	}
	fmt.Println("---------------------------------")

}
