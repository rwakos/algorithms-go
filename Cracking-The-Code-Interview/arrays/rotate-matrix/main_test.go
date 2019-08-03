package main

import (
	"fmt"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	m1 := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}}
	m1Rotated := [][]string{{"g", "d", "a"}, {"h", "e", "b"}, {"i", "f", "c"}}

	m2 := [][]string{{"a", "b", "c", "d"}, {"e", "f", "g", "h"}, {"i", "j", "k", "l"}, {"m", "n", "o", "p"}}
	m2Rotated := [][]string{{"m", "i", "e", "a"}, {"n", "j", "f", "b"}, {"o", "k", "g", "c"}, {"p", "l", "h", "d"}}

	if !compareMatrix(rotateMatrix(m1), m1Rotated) {
		t.Errorf("FAIL Matrix 3x3")
	} else {
		fmt.Println("PASS - Case 1")
	}

	if !compareMatrix(rotateMatrix(m2), m2Rotated) {
		t.Errorf("FAIL Matrix 4x4")
		printMatrix(m2)
		printMatrix(rotateMatrix(m2))
	} else {
		fmt.Println("PASS - Case 2")
	}
}
