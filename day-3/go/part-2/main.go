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

type Row []rune
type Grid []Row

type Gear struct {
	RowIndex    int
	Index       int
	GearNumbers []int
}

type Number struct {
	Value      int
	RowIndex   int
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
	numbers := findNumbersInGrid(grid)
	gears := findGears(grid, numbers)

	for _, gear := range gears {
		sum += (gear.GearNumbers[0] * gear.GearNumbers[1])
	}
	fmt.Println(sum)

}

func parseSchematic(es []string) Grid {
	grid := make(Grid, len(es))
	for i, line := range es {
		grid[i] = Row(line)
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

func findNumbersInGrid(grid Grid) []Number {
	var numbers []Number
	for rowIndex, row := range grid {
		numbers = append(numbers, findNumbersInRow(rowIndex, row)...)
	}
	return numbers
}

func findNumbersInRow(rowIndex int, row Row) []Number {
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
				numbers = append(numbers, createNumber(current, rowIndex, startIndex, i-1))
				current = []rune{}
			}
			startIndex = -1
		}
	}

	if len(current) > 0 {
		numbers = append(numbers, createNumber(current, rowIndex, startIndex, len(row)-1))
	}

	return numbers
}

func createNumber(runes []rune, rowIndex int, startIndex, endIndex int) Number {
	number, err := strconv.Atoi(string(runes))
	if err != nil {
		fmt.Println(err)
	}
	return Number{Value: number, RowIndex: rowIndex, StartIndex: startIndex, EndIndex: endIndex}
}

func findGears(grid Grid, numbers []Number) []Gear {
	var gears []Gear
	for rowIndex, row := range grid {
		for symbolIndex, symbol := range row {
			if symbol == '*' {
				gearNumbers := getGearNumbers(rowIndex, symbolIndex, numbers)
				if len(gearNumbers) == 2 {
					gears = append(gears, Gear{RowIndex: rowIndex, Index: symbolIndex, GearNumbers: gearNumbers})
				}
			}
		}
	}
	return gears
}

func getGearNumbers(gearRowIndex int, gearIndex int, numbers []Number) []int {
	var gearNumbers []int
	for _, number := range numbers {
		if number.StartIndex <= gearIndex+1 && number.EndIndex >= gearIndex-1 && (number.RowIndex >= gearRowIndex-1 && number.RowIndex <= gearRowIndex+1) {
			gearNumbers = append(gearNumbers, number.Value)
		}
	}
	return gearNumbers
}
