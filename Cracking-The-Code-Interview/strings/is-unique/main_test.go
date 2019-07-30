package main

import (
	"fmt"
	"testing"
)

func TestHasUniquesCharacters(t *testing.T) {
	good := []string{"abcdefx", "bdgrotny", "zxcvbnmlkjhgfdsaqwertyu"}
	bad := []string{"abcdefa", "bb", "asdws"}

	for _, v := range good {
		if !hasUniquesCharacters(v) {
			t.Errorf("This is not a Unique string: %s", v)
		} else {
			fmt.Println("GOOD: ", v)
		}
	}

	for _, v := range bad {
		if hasUniquesCharacters(v) {
			t.Errorf("This is a Unique string: %s", v)
		} else {
			fmt.Println("GOOD_FAIL: ", v)
		}
	}
}
