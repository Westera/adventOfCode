package main

import (
	"adventOfCode/2020/util"
	"fmt"
	"log"
	"os"
	"sort"
)

func part1(expenses []int, target int) int{
	sort.Ints(expenses)
	pointerLow, pointerHight := 0, len(expenses) - 1
	sum := expenses[pointerLow] + expenses[pointerHight]

	for sum != target {
		if sum > target {
			pointerHight--
		} else {
			pointerLow++
		}
		if pointerHight == pointerLow {
			return -1
		}
		sum = expenses[pointerLow] + expenses[pointerHight]
	}

	return expenses[pointerLow] * expenses [pointerHight]
}

func part2(expenses []int) int{
	sort.Ints(expenses)
	for pointer := 0 ; pointer < len(expenses) - 2 ; pointer++ {
		pointerValue := expenses[pointer]
		partialResult := part1(expenses[pointer + 1 : len(expenses) - 1], 2020 - pointerValue)
		if partialResult != -1 {
			return partialResult * pointerValue
		}
	}
	return -1
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	intArray, err := util.ReadIntsToArray(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(intArray, 2020)
	resultPt2 := part2(intArray)

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
