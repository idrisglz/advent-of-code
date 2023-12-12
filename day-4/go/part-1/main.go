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
	cards, err := readInputFile(filename)
	if err != nil {
		fmt.Printf("unable to read the input file: %v. exiting\n", err)
		return
	}

	sum := calculateTotal(cards)

	fmt.Println(sum)
}

func calculateTotal(cards []string) int {
	var sum int
	for _, card := range cards {
		winners, actualNumbers, err := parseCard(card)
		if err != nil {
			fmt.Println("unable to extract numbers, skipping")
			continue
		}

		winnerCount := countWinners(winners, actualNumbers)

		sum += int(math.Pow(2, float64(winnerCount)-1))
	}
	return sum
}

func parseCard(card string) (winners []int, actualNumbers []int, err error) {
	parts := strings.Split(card, "|")
	if len(parts) != 2 {
		return nil, nil, fmt.Errorf("invalid card format")
	}

	winners, err = extractNumbers(strings.Split(parts[0], ":")[1])
	if err != nil {
		return nil, nil, err
	}

	actualNumbers, err = extractNumbers(parts[1])
	if err != nil {
		return nil, nil, err
	}

	return winners, actualNumbers, nil
}

func extractNumbers(s string) ([]int, error) {
	fmt.Println(s)
	var numbers []int
	for _, str := range strings.Fields(s) {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %v", err)
		}
		numbers = append(numbers, num)
	}
	return numbers, nil
}

func countWinners(winners []int, actualNumbers []int) int {
	var winnerCount int
	for _, actualNumber := range actualNumbers {
		if slices.Contains(winners, actualNumber) {
			winnerCount++
		}
	}
	return winnerCount
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
