package main

import (
	"fmt"
	"testing"
)

func TestSmallestString(t *testing.T) {
	if smallestString64(25) != "ACC" {
		t.Errorf("The answer should be ACC")
	} else {
		fmt.Println("Good")
	}

	if smallestString64(64) != "ABCCCCC" {
		t.Errorf("The answer should be ABCCCCC")
	} else {
		fmt.Println("Good")
	}
}
