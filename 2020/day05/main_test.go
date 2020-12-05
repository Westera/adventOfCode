package main

import "testing"

func TestPart1(t *testing.T) {
	actual := part1([]string{"BFFFBBFRRR", "FFFBBBFRRR", "BBFFBBFRLL"})
	if actual != 820 {
		t.Errorf("Expected %d to be 820", actual)
	}
}