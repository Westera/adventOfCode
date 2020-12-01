package main

import "testing"

func TestPart1(t *testing.T) {
	actual := part1([]int{1721, 979, 366, 299, 675, 1456}, 2020)
	if actual != 514579 {
		t.Errorf("Expected %d to be 514579", actual)
	}
}

func TestPart2(t *testing.T) {
	actual := part2([]int{1721, 979, 366, 299, 675, 1456})
	if actual != 241861950 {
		t.Errorf("Expected %d to be 241861950", actual)
	}
}
