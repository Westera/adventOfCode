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

type Passport struct {
	byr bool
	iyr bool
	eyr bool
	hgt bool
	hcl bool
	ecl bool
	pid bool
}

func part1(batchPassports []string, shouldValidate bool) int {
	nrOfValid := 0
	rowRegexp := regexp.MustCompile("[ :]")
	passport := newPassport()
	for _, passportRow := range batchPassports {
		if passportRow == "" {
			if isValidPassport(passport) {
				nrOfValid++
			}
			passport = newPassport()
		} else {
			passportData := rowRegexp.Split(passportRow, -1)
			for i := 0; i < len(passportData); i = i + 2 {
				passport = validatePassportData(passport, passportData[i], passportData[i+1], shouldValidate)
			}
		}

	}
	if isValidPassport(passport) {
		nrOfValid++
	}
	return nrOfValid
}

func part2(batchPassports []string) int {
	return part1(batchPassports, true)
}

func newPassport() Passport {
	return Passport{
		byr: false,
		iyr: false,
		eyr: false,
		hgt: false,
		hcl: false,
		ecl: false,
		pid: false,
	}
}

func validatePassportData(passport Passport, field string, data string, shouldValidate bool) Passport {
	hairColorValidator := regexp.MustCompile("^#[[:xdigit:]]{6}$")
	eyeColorOptions := "amb blu brn gry grn hzl oth"
	pidValidator := regexp.MustCompile("^[0-9]{9}$")

	switch field {
	case "byr":
		converted, _ := strconv.Atoi(data)
		passport.byr = !shouldValidate || (1920 <= converted && converted <= 2002)
	case "iyr":
		converted, _ := strconv.Atoi(data)
		passport.iyr = !shouldValidate || (2010 <= converted && converted <= 2020)
	case "eyr":
		converted, _ := strconv.Atoi(data)
		passport.eyr = !shouldValidate || (2020 <= converted && converted <= 2030)
	case "hgt":
		if shouldValidate {
			unit := data[len(data)-2:]
			if unit == "cm" {
				height, err := strconv.Atoi(data[0:3])
				if err == nil {
					passport.hgt = 150 <= height && height <= 193
				}
			} else if unit == "in" {
				height, err := strconv.Atoi(data[0:2])
				if err == nil {
					passport.hgt = 59 <= height && height <= 76
				}
			}
		} else {
			passport.hgt = true
		}
	case "hcl":
		passport.hcl = !shouldValidate || hairColorValidator.MatchString(data)
	case "ecl":
		passport.ecl = !shouldValidate || strings.Contains(eyeColorOptions, data)
	case "pid":
		passport.pid = !shouldValidate || pidValidator.MatchString(data)
	}
	return passport
}

func isValidPassport(passport Passport) bool {
	return passport.byr && passport.iyr && passport.eyr && passport.hgt && passport.hcl && passport.ecl && passport.pid
}

func main() {
	file, err := os.Open("2020/day04/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	lines, err := util.ReadToArray(file)

	if err != nil {
		log.Fatal(err)
	}

	resultPt1 := part1(lines, false)
	resultPt2 := part2(lines)

	fmt.Println(resultPt1)
	fmt.Println(resultPt2)
}
