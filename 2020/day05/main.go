package main

import (
	"adventOfCode/2020/util"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
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
	seatIds := []int{len(boardingPasses)}
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
	replacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	binaryBoardingPass := replacer.Replace(boardingPass)
	row, _ := strconv.ParseInt(binaryBoardingPass[0:7], 2, 16)
	col, _ := strconv.ParseInt(binaryBoardingPass[7:], 2, 16)
	return int(row * 8 + col)
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
