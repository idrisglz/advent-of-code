package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const filename = "../input.txt"

func main() {
	input := readInputFile(filename)
	var sum int

	for _, card := range input {
		fmt.Println(card)
		parts := strings.Split(card, "|")
		winnersPart := strings.Split(parts[0], ":")
		winnersPart = strings.Split(winnersPart[1], " ")

		actualsPart := strings.Split(parts[1], " ")
		winners := []int{}
		actualNumbers := []int{}
		var winnerCount int

		for _, winner := range winnersPart {
			if winner != "" {
				num, err := strconv.Atoi(strings.TrimSpace(winner))
				if err != nil {
					fmt.Println("Unable to convert to a number")
				}
				winners = append(winners, num)
			}
		}

		for _, actual := range actualsPart {
			if actual != "" {
				num, err := strconv.Atoi(strings.TrimSpace(actual))
				if err != nil {
					fmt.Println("Unable to convert to a number")
				}
				actualNumbers = append(actualNumbers, num)
			}

		}

		for _, actualNumber := range actualNumbers {
			if slices.Contains(winners, actualNumber) {
				winnerCount++
			}
		}

		result := int(math.Pow(2, float64(winnerCount)-1))
		fmt.Println(result)
		sum += result
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
