package main

import (
	"fmt"
	"testing"
)

func TestIsPermutation(t *testing.T) {
	if !isPermutation("richard", "drahcir") {
		t.Errorf("This is a permutation, check the code")
	} else {
		fmt.Println("Good")
	}

	if isPermutation("abcd", "fghi") {
		t.Errorf("This shouldn't be a permutation")
	} else {
		fmt.Println("Good")
	}

	if isPermutation("abcd", "cd") {
		t.Errorf("This shouldn't be a permutation")
	} else {
		fmt.Println("Good")
	}
}
