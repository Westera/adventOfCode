package util

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

func ReadIntsToArray(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func ReadToArray(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func ReadToString(r io.Reader) (string, error) {
	content, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	// Convert []byte to string and print to screen
	text := string(content)
	return text, nil
}