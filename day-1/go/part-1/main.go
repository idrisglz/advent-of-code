package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

const (
	filename = "../input.txt"
)

func main() {
	inputs := readInputFile(filename)
	fmt.Println(calculateSum(inputs))
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

func calculateSum(values []string) int {
	sum := 0
	for _, text := range values {
		num, err := getCalibrationValues(text)
		if err != nil {
			fmt.Println(err)
			return -1
		}
		sum += num
	}
	return sum
}

func getCalibrationValues(input string) (int, error) {
	var firstDigit, lastDigit string

	for _, r := range input {
		if unicode.IsDigit(r) {
			firstDigit = string(r)
			break
		}
	}

	for i := len(input); i > 0; i-- {
		if unicode.IsDigit(rune(input[i-1])) {
			lastDigit = string(input[i-1])
			break
		}

	}

	return strconv.Atoi(firstDigit + lastDigit)
}
