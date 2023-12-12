package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	filename = "../input.txt"
)

var (
	greenRegex = regexp.MustCompile(`(\d+)\s*green`)
	blueRegex  = regexp.MustCompile(`(\d+)\s*blue`)
	redRegex   = regexp.MustCompile(`(\d+)\s*red`)
)

func main() {
	inputs, err := readInputFile(filename)
	if err != nil {
		fmt.Println("unable to read the input file. exiting")
		return
	}
	var sum int
	for _, input := range inputs {
		green, blue, red := getMaxCount(input, greenRegex, blueRegex, redRegex)
		power := green * blue * red
		sum += power
	}

	fmt.Println(sum)
}

func readInputFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getMaxCount(input string, greenRegex *regexp.Regexp, blueRegex *regexp.Regexp, redRegex *regexp.Regexp) (int, int, int) {
	maxGreen := getMax(input, greenRegex)
	maxBlue := getMax(input, blueRegex)
	maxRed := getMax(input, redRegex)

	return maxGreen, maxBlue, maxRed
}

func getMax(input string, colorRegex *regexp.Regexp) int {
	var maxCount int

	matches := colorRegex.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		count, err := strconv.Atoi(match[1])

		if err != nil {
			fmt.Println(err)
			continue
		}
		maxCount = max(maxCount, count)
	}
	return maxCount
}
