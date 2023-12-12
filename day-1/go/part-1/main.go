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
	inputs, err := readInputFile(filename)
	if err != nil {
		fmt.Printf("unable to read the input file: %v. exiting\n", err)
		return
	}
	fmt.Println(calculateSum(inputs))
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
