package main

import (
	"adventOfCode/2020/util"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func part1(passwords []string) int{
	valid := 0
	rowRegexp := regexp.MustCompile(" |-|:")
	for _, password := range passwords {
		splitPassword := rowRegexp.Split(password, -1)
		min, _ := strconv.Atoi(splitPassword[0])
		max, _ := strconv.Atoi(splitPassword[1])
		if(validatePassword(min, max, splitPassword[2], splitPassword[4])) {
			valid++
		}
	}
	return valid
}

func part2(passwords []string) int{
	valid := 0
	rowRegexp := regexp.MustCompile(" |-|:")
	for _, password := range passwords {
		splitPassword := rowRegexp.Split(password, -1)
		index1, _ := strconv.Atoi(splitPassword[0])
		index2, _ := strconv.Atoi(splitPassword[1])
		if(validatePasswordV2(index1 - 1, index2 - 1, splitPassword[2], splitPassword[4])) {
			valid++
		}
	}
	return valid
}

func validatePassword(minOccurrence int, maxOccurrence int, search string, password string) bool{
	occurrence := strings.Count(password, search)
	return minOccurrence <= occurrence && occurrence <= maxOccurrence
}

func validatePasswordV2(index1 int, index2 int, search string, password string) bool{
	return (password[index1] == search[0]) != (password[index2] == search[0])
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
