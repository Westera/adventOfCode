package main

import "testing"

func TestPart1(t *testing.T) {
	actual := part1([]string{"1-3 a: abcde",
		                     "1-3 b: cdefg",
		                     "2-9 c: ccccccccc"})
	if actual != 2 {
		t.Errorf("Expected %d to be 2", actual)
	}
}

func TestPart2(t *testing.T) {
	actual := part2([]string{"1-3 a: abcde",
		                     "1-3 b: cdefg",
		                     "2-9 c: ccccccccc"})
	if actual != 1 {
		t.Errorf("Expected %d to be 1", actual)
	}
}
