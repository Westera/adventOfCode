package main

import "testing"

func TestPart1(t *testing.T) {
	actual := part1([]string{"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b"})
	if actual != 11 {
		t.Errorf("Expected %d to be 11", actual)
	}
}

func TestPart2(t *testing.T) {
	actual := part2([]string{"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b"})
	if actual != 6 {
		t.Errorf("Expected %d to be 6", actual)
	}
}
