package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	digitTexts := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	var firstDigit, lastDigit string

first_digit_loop:
	for i, r := range input {
		if unicode.IsDigit(r) {
			firstDigit = string(r)
			break
		}

		for text, digit := range digitTexts {
			if strings.HasPrefix(input[i:], text) {
				firstDigit = digit
				break first_digit_loop
			}
		}
	}

last_digit_loop:
	for i := len(input); i > 0; i-- {
		if unicode.IsDigit(rune(input[i-1])) {
			lastDigit = string(input[i-1])
			break
		}
		for text, digit := range digitTexts {
			if strings.HasSuffix(input[:i], text) {
				lastDigit = digit
				break last_digit_loop
			}
		}
	}

	return strconv.Atoi(firstDigit + lastDigit)
}
