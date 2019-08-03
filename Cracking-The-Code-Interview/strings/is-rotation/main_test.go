package main

import (
	"fmt"
	"testing"
)

func TestIsRotation(t *testing.T) {
	if !isRotation("water", "terwa") {
		t.Errorf("FAIL, this should be a rotation")
	} else {
		fmt.Println("PASS - Case 1")
	}

	if !isRotation("thewateriswhite", "iswhitethewater") {
		t.Errorf("FAIL, this should be a rotation")
	} else {
		fmt.Println("PASS - Case 2")
	}

	if isRotation("sometest", "testnew") {
		t.Errorf("FAIL, this should'nt be a rotation")
	} else {
		fmt.Println("PASS - Case 3")
	}
}
