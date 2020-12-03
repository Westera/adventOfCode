package main

import (
	"adventOfCode/2020/util"
	"fmt"
	"log"
	"os"
)

type Point struct {
	X int
	Y int
}

func part1(slopeMap []string, x int, y int) int{
	nrOfTrees := 0
	width := len(slopeMap[0])
	currentLocation := Point{0,0}

	for currentLocation.Y < len(slopeMap) {
		if slopeMap[currentLocation.Y][currentLocation.X] == '#' {
			nrOfTrees++
		}
		currentLocation = Point{(currentLocation.X + x)% width, currentLocation.Y + y}
	}

	return nrOfTrees
}

func part2(slopeMap []string) int{
	result := part1(slopeMap, 1, 1)
	result *= part1(slopeMap, 3, 1)
	result *= part1(slopeMap, 5, 1)
	result *= part1(slopeMap, 7, 1)
	result *= part1(slopeMap, 1, 2)
	return result
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines, err := util.ReadToArray(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(lines, 3, 1)
	resultPt2 := part2(lines)

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
