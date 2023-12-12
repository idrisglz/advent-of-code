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

type EngineSchematic []string
type Row []rune
type Grid []Row

type Number struct {
	Value      int
	StartIndex int
	EndIndex   int
}

func main() {
	input, err := readInputFile(filename)
	if err != nil {
		fmt.Println("unable to read the input file. exiting")
		return
	}

	sum := 0
	grid := parseSchematic(input)
	eligibilityGrid := initEligibilityGrid(len(grid), len(grid[0]))

	for rowIndex, row := range grid {
		for runeIndex, r := range row {
			if !unicode.IsDigit(r) && r != '.' {
				eligibilityGrid[rowIndex][runeIndex] = '+'
				eligibilityGrid[rowIndex][runeIndex-1] = '+'
				eligibilityGrid[rowIndex][runeIndex+1] = '+'

				if rowIndex != 0 {
					eligibilityGrid[rowIndex-1][runeIndex] = '+'
					eligibilityGrid[rowIndex-1][runeIndex-1] = '+'
					eligibilityGrid[rowIndex-1][runeIndex+1] = '+'
				}

				if rowIndex < len(grid)-1 {
					eligibilityGrid[rowIndex+1][runeIndex] = '+'
					eligibilityGrid[rowIndex+1][runeIndex-1] = '+'
					eligibilityGrid[rowIndex+1][runeIndex+1] = '+'
				}
			}
		}
	}

	for rowIndex, row := range grid {
		numbers := findNumbers(row)

	number_loop:
		for _, number := range numbers {
			for i := number.StartIndex; i <= number.EndIndex; i++ {
				if eligibilityGrid[rowIndex][i] == '+' {
					sum += number.Value
					continue number_loop
				}
			}
		}
	}
	fmt.Println(sum)
}

func parseSchematic(es EngineSchematic) Grid {
	grid := make(Grid, len(es))
	for i, line := range es {
		grid[i] = Row(line)
	}
	return grid
}

func initEligibilityGrid(cols, rows int) Grid {
	grid := make(Grid, rows)

	for i := 0; i < cols; i++ {
		grid[i] = make(Row, rows)
	}
	return grid
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

func findNumbers(input []rune) []Number {
	numbers := []Number{}
	var current []rune
	startIndex := -1
	lastIndex := len(input) - 1
	for i, r := range input {
		if unicode.IsDigit(r) {
			current = append(current, r)
			if startIndex == -1 {
				startIndex = i
			}
		} else if len(current) > 0 {
			num := createNumber(current, startIndex, i-1)
			numbers = append(numbers, num)
			current = []rune{}
			startIndex = -1
		}
	}

	if len(current) > 0 {
		num := createNumber(current, startIndex, lastIndex)
		numbers = append(numbers, num)
	}

	return numbers
}

func createNumber(runes []rune, start, end int) Number {
	num, err := strconv.Atoi(string(runes))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse number: %s", string(runes)))
	}
	return Number{Value: num, StartIndex: start, EndIndex: end}
}
