package main

import (
	"fmt"
	"testing"
)

func TestOneEditAway(t *testing.T) {
	if !oneEditAway("pale", "ple") {
		t.Errorf("FAIL: this should be true")
	} else {
		fmt.Println("PASS - Case 1")
	}

	if !oneEditAway("pale", "bale") {
		t.Errorf("FAIL: this should be true")
	} else {
		fmt.Println("PASS - Case 2")
	}

	if oneEditAway("pale", "bae") {
		t.Errorf("FAIL: this should be false")
	} else {
		fmt.Println("PASS - Case 3")
	}
}
