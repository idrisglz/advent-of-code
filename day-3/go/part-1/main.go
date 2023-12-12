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
		fmt.Printf("unable to read the input file: %v. exiting\n", err)
		return
	}

	var sum int
	grid := parseSchematic(input)
	eligibilityGrid := initEligibilityGrid(len(grid), len(grid[0]))

	markEligibility(grid, eligibilityGrid)

	sum = calculateSum(grid, eligibilityGrid)
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

func markEligibility(grid Grid, eligibilityGrid Grid) {
	for rowIndex, row := range grid {
		for runeIndex, r := range row {
			if !unicode.IsDigit(r) && r != '.' {
				updateEligibilityGrid(eligibilityGrid, rowIndex, runeIndex, len(grid), len(row))
			}
		}
	}
}

func updateEligibilityGrid(grid Grid, rowIndex, runeIndex, maxRow, maxCol int) {
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newRow, newCol := rowIndex+i, runeIndex+j
			if newRow >= 0 && newRow < maxRow && newCol >= 0 && newCol < maxCol {
				grid[newRow][newCol] = '+'
			}
		}
	}
}

func isEligible(number Number, rowIndex int, eligibilityGrid Grid) bool {
	for i := number.StartIndex; i <= number.EndIndex; i++ {
		if eligibilityGrid[rowIndex][i] == '+' {
			return true
		}
	}
	return false
}

func calculateSum(grid Grid, eligibilityGrid Grid) int {
	var sum int
	for rowIndex, row := range grid {
		numbers := findNumbers(row)
		for _, number := range numbers {
			if isEligible(number, rowIndex, eligibilityGrid) {
				sum += number.Value
			}
		}
	}
	return sum
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

func findNumbers(row Row) []Number {
	var numbers []Number
	var current []rune
	startIndex := -1

	for i, r := range row {
		if unicode.IsDigit(r) {
			if startIndex == -1 {
				startIndex = i
			}
			current = append(current, r)
		} else {
			if len(current) > 0 {
				numbers = append(numbers, createNumber(current, startIndex, i-1))
				current = []rune{}
			}
			startIndex = -1
		}
	}

	if len(current) > 0 {
		numbers = append(numbers, createNumber(current, startIndex, len(row)-1))
	}

	return numbers
}

func createNumber(runes []rune, startIndex, endIndex int) Number {
	number, err := strconv.Atoi(string(runes))
	if err != nil {
		fmt.Println(err)
	}
	return Number{Value: number, StartIndex: startIndex, EndIndex: endIndex}
}
