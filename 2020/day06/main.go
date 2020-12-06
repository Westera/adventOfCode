package main

import (
	"adventOfCode/util"
	"fmt"
	"log"
	"os"
)

func part1(customDeclarations []string) int {
	numberOfQuestionsAnswered := 0
	answerMap := make(map[int32]bool)
	for _, customDeclaration := range customDeclarations {
		if customDeclaration == "" {
			numberOfQuestionsAnswered += len(answerMap)
			answerMap = make(map[int32]bool)
		} else {
			for _, answer := range customDeclaration {
				answerMap[answer] = true
			}
		}
	}
	numberOfQuestionsAnswered += len(answerMap)

	return numberOfQuestionsAnswered
}

func part2(customDeclarations []string) int {
	numberOfQuestionsAnswered := 0
	numberInGroup := 0
	answerMap := make(map[int32]int)
	for _, customDeclaration := range customDeclarations {
		if customDeclaration == "" {
			for key := range answerMap {
				if answerMap[key] == numberInGroup {
					numberOfQuestionsAnswered++
				}
			}
			numberInGroup = 0
			answerMap = make(map[int32]int)
		} else {
			numberInGroup++
			for _, answer := range customDeclaration {
				answerMap[answer]++
			}
		}
	}
	for key := range answerMap {
		if answerMap[key] == numberInGroup {
			numberOfQuestionsAnswered++
		}
	}

	return numberOfQuestionsAnswered
}

func main() {
	file, err := os.Open("2020/day06/input.txt")

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
