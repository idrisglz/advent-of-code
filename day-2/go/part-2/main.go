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
	sum := 0
	inputs := readInputFile(filename)
	for _, input := range inputs {
		green, blue, red := getMaxCount(input, greenRegex, blueRegex, redRegex)
		power := green * blue * red
		sum += power
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
