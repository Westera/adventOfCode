package main

import "testing"

func TestPart1(t *testing.T) {
	actual := part1([]string{"..##.......",
		                     "#...#...#..",
		                     ".#....#..#.",
		                     "..#.#...#.#",
		                     ".#...##..#.",
		                     "..#.##.....",
		                     ".#.#.#....#",
		                     ".#........#",
		                     "#.##...#...",
		                     "#...##....#",
		                     ".#..#...#.#"}, 3,1)
	if actual != 7 {
		t.Errorf("Expected %d to be 7", actual)
	}
}

func TestPart2(t *testing.T) {
	actual := part2([]string{"..##.......",
		                     "#...#...#..",
		                     ".#....#..#.",
	 	                     "..#.#...#.#",
		                     ".#...##..#.",
		                     "..#.##.....",
		                     ".#.#.#....#",
		                     ".#........#",
		                     "#.##...#...",
		                     "#...##....#",
		                     ".#..#...#.#"})
	if actual != 336 {
		t.Errorf("Expected %d to be 336", actual)
	}
}
