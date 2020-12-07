package main

import (
	"adventOfCode/util"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type InsideBag struct {
	bagColor string
	amount int
}

func part1(bagRules []string, target string) int {
	removeExtraInfo := regexp.MustCompile("[0-9] |contain| bag[s]?[\\. ]?|,")
	whatCanBeInWhatMap := make(map[string][]string, len(bagRules))
	for _, bagRule := range bagRules {
		bags := strings.Split(removeExtraInfo.ReplaceAllString(bagRule, ""), " ")
		outsideBag := fmt.Sprintf("%s %s", bags[0], bags[1])
		for i := 2 ; i < len(bags) ; i += 2 {
			insideBag := fmt.Sprintf("%s %s", bags[i], bags[i+1])
			currentBags := whatCanBeInWhatMap[insideBag]
			whatCanBeInWhatMap[insideBag] = append(currentBags, outsideBag)
		}
	}
	return findNumberOfBags(whatCanBeInWhatMap, make(map[string]bool, len(bagRules)), target)
}

func part2(bagRules []string, target string) int {
	removeExtraInfo := regexp.MustCompile("contain| bag[s]?[\\. ]?|,")
	whatIsInWhatMap := make(map[string][]InsideBag, len(bagRules))
	for _, bagRule := range bagRules {
		bags := strings.Split(removeExtraInfo.ReplaceAllString(bagRule, ""), " ")
		outsideBag := fmt.Sprintf("%s %s", bags[0], bags[1])
		for i := 2 ; i < len(bags) ; i += 3 {
			if bags[i] != "no" {
				nrOfBags, _ := strconv.Atoi(bags[i])
				insideBag := fmt.Sprintf("%s %s", bags[i+1], bags[i+2])
				currentBags := whatIsInWhatMap[outsideBag]
				whatIsInWhatMap[outsideBag] = append(currentBags, InsideBag{insideBag, nrOfBags})
			} else {
				i -= 1
			}
		}
	}
	return howManyBagsToBring(whatIsInWhatMap, target)
}

func findNumberOfBags(whatCanBeInWhatMap map[string][]string, whatHasBeenSeen map[string]bool, target string) int{
	if len(whatCanBeInWhatMap[target]) != 0 {
		for _, bag := range whatCanBeInWhatMap[target] {
			whatHasBeenSeen[bag] = true
			findNumberOfBags(whatCanBeInWhatMap,whatHasBeenSeen, bag)
		}
	}
	return len(whatHasBeenSeen)
}

func howManyBagsToBring(whatIsInWhatMap map[string][]InsideBag, target string) int{
	if len(whatIsInWhatMap[target]) == 0 {
		return 0
	} else {
		sum := 0
		for _, insideBag := range whatIsInWhatMap[target] {
			sum += insideBag.amount
			sum += insideBag.amount*howManyBagsToBring(whatIsInWhatMap, insideBag.bagColor)
		}
		return sum
	}
}

func main() {
	file, err := os.Open("2020/day07/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines, err := util.ReadToArray(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(lines, "shiny gold")
	resultPt2 := part2(lines, "shiny gold")

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
