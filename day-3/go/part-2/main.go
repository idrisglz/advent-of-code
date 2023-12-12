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
	input := readInputFile(filename)
	numbers := []Number{}

	sum := 0
	grid := parseSchematic(input)
	// eligibilityGrid := initEligibilityGrid(len(grid), len(grid[0]))

	for rowIndex, row := range grid {
		numbers = append(numbers, findNumbers(rowIndex, row)...)
	}

	gears := findGears(grid, numbers)

	for _, gear := range gears {
		sum += (gear.GearNumbers[0] * gear.GearNumbers[1])
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

func readInputFile(filename string) EngineSchematic {
	file, err := os.Open(filename)
	if err != nil {
		panic("Unable to read the input file.")
	}

	defer file.Close()

	var es EngineSchematic
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		es = append(es, scanner.Text())
	}
	return es
}

func findNumbers(rowIndex int, row []rune) (numbers []Number) {
	var current []rune
	startIndex := -1
	lastIndex := len(row) - 1
	for i, r := range row {
		if unicode.IsDigit(r) {
			current = append(current, r)
			if startIndex == -1 {
				startIndex = i
			}
		} else if len(current) > 0 {
			num := createNumber(current, rowIndex, startIndex, i-1)
			numbers = append(numbers, num)
			current = []rune{}
			startIndex = -1
		}
	}

	if len(current) > 0 {
		num := createNumber(current, rowIndex, startIndex, lastIndex)
		numbers = append(numbers, num)
	}

	return numbers
}

func createNumber(runes []rune, rowIndex int, startIndex, endIndex int) Number {
	number, err := strconv.Atoi(string(runes))
	if err != nil {
		panic(fmt.Sprintf("Failed to parse number: %s", string(runes)))
	}
	return Number{Value: number, RowIndex: rowIndex, StartIndex: startIndex, EndIndex: endIndex}
}

func findGears(es Grid, numbers []Number) (gears []Gear) {
	for rowIndex, row := range es {
		for symbolIndex, symbol := range row {
			if symbol == '*' && len(getGearNumbers(rowIndex, symbolIndex, numbers)) == 2 {
				gears = append(gears, createGear(rowIndex, symbolIndex, getGearNumbers(rowIndex, symbolIndex, numbers)))
			}
		}
	}
	return gears
}

func createGear(rowIndex, index int, gearNums []int) Gear {
	return Gear{RowIndex: rowIndex, Index: index, GearNumbers: gearNums}
}

func getGearNumbers(gearRowIndex int, gearIndex int, numbers []Number) (gearNumbers []int) {
	for _, number := range numbers {
		if number.StartIndex <= gearIndex+1 && number.EndIndex >= gearIndex-1 && (number.RowIndex > gearRowIndex-2 && number.RowIndex < gearRowIndex+2) {
			gearNumbers = append(gearNumbers, number.Value)
		}
	}
	return gearNumbers
}
