package main

import (
	"adventOfCode/2020/util"
	"fmt"
	"log"
	"os"
	"sort"
)

func part1(boardingPasses []string) int{
	highestSeatID := 0

	for _, boardingPass := range boardingPasses {
		seatId := calculateSeatId(boardingPass)
		if seatId > highestSeatID {
			highestSeatID = seatId
		}
	}

	return highestSeatID
}

func part2(boardingPasses []string) int{
	seatIds := []int{}
	for _, boardingPass := range boardingPasses {
		seatId := calculateSeatId(boardingPass)
		seatIds = append(seatIds, seatId)
	}
	sort.Ints(seatIds)

	for i, seatId := range seatIds {
		if seatId + 1 != seatIds[i+1]{
			return seatId + 1
		}
	}

	log.Fatal("Seat not found")
	return -1
}

func calculateSeatId(boardingPass string) int{
	distanceRow, distanceColumn, minRow, maxRow, minCol, maxCol := 128, 8, 0, 127, 0, 7
	for _, char := range boardingPass {
		switch char {
		case 'F':
			distanceRow /= 2
			maxRow -= distanceRow
		case 'B':
			distanceRow /= 2
			minRow += distanceRow
		case 'L':
			distanceColumn /= 2
			maxCol -= distanceColumn
		case 'R':
			distanceColumn /= 2
			minCol += distanceColumn
		}
	}
	return minRow * 8 + minCol
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

	resultPt1 := part1(lines)
	resultPt2 := part2(lines)

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
