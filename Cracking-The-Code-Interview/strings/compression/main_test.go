package main

import (
	"fmt"
	"testing"
)

func TestCompressString(t *testing.T) {
	result := compressString("aaabbb")
	if result == "a3b2" {
		t.Errorf("FAIL: this should be a3b2, not %s", result)
	} else {
		fmt.Println("PASS - Case 1")
	}

	result = compressString("abcd")
	if result != "abcd" {
		t.Errorf("FAIL: this should be abcd, not %s", result)
	} else {
		fmt.Println("PASS - Case 2")
	}

	result = compressString("abbbd")
	if result != "ab3d" {
		t.Errorf("FAIL: this should be ab3d, not %s", result)
	} else {
		fmt.Println("PASS - Case 3")
	}
}
