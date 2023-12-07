package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
	filename   = "../input.txt"
)

var (
	greenRegex = regexp.MustCompile(`(\d+)\s*green`)
	blueRegex  = regexp.MustCompile(`(\d+)\s*blue`)
	redRegex   = regexp.MustCompile(`(\d+)\s*red`)
)

func main() {
	var sum int

	inputs := readInputFile(filename)

	for _, input := range inputs {
		id, err := parseGameId(input)
		if err != nil {
			panic(err)
		}

		if isEligible(input, greenRegex, blueRegex, redRegex) {
			sum += id
		}
	}
	fmt.Println(sum)
}

func readInputFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic("Unable to read the input file")
	}

	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs
}

func isEligible(input string, greenRegex *regexp.Regexp, blueRegex *regexp.Regexp, redRegex *regexp.Regexp) bool {
	startIndex := strings.Index(input, ":")
	colorValues := input[startIndex+1:]
	parts := strings.Split(colorValues, ";")

	for _, p := range parts {
		green, greenErr := getCount(p, greenRegex)
		blue, blueErr := getCount(p, blueRegex)
		red, redErr := getCount(p, redRegex)

		if greenErr != nil || blueErr != nil || redErr != nil {
			panic("Error occurred while getting count")
		}

		if green > greenLimit || red > redLimit || blue > blueLimit {
			return false
		}
	}
	return true
}

func parseGameId(input string) (id int, err error) {
	spaceIndex := strings.Index(input, " ")
	colonIndex := strings.Index(input, ":")

	if spaceIndex == -1 || colonIndex == -1 || spaceIndex >= colonIndex {
		return 0, fmt.Errorf("invalid format")
	}

	idStr := input[spaceIndex+1 : colonIndex]
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %w", err)
	}

	return id, nil
}

func getCount(input string, colorRegex *regexp.Regexp) (count int, err error) {
	match := colorRegex.FindStringSubmatch(input)
	if match == nil {
		return 0, nil
	}

	count, err = strconv.Atoi(match[1])
	if err != nil {
		return 0, fmt.Errorf("invalid count: %w", err)
	}

	return count, nil
}
