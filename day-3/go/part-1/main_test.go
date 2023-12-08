package main

import (
	"testing"
)

// var testCases = []struct {
// 	input          []string
// 	expectedResult int
// }{
// 	{
// 		input: []string{
// 			"467..114..",
// 			"...*......",
// 			"..35..633.",
// 			"......#...",
// 			"617*......",
// 			".....+.58.",
// 			"..592.....",
// 			"......755.",
// 			"...$.*....",
// 			".664.598..",
// 		},
// 		expectedResult: 4361,
// 	},
// }

var numberTestCases = []struct {
	input          []rune
	expectedResult []Number
}{
	{[]rune("467..114.."), []Number{{467, 0, 2}, {114, 5, 7}}},
	{[]rune("...*......"), []Number{}},
	{[]rune("..35..633."), []Number{{35, 2, 3}, {633, 6, 8}}},
	{[]rune("......#..."), []Number{}},
	{[]rune("617*......"), []Number{{617, 0, 2}}},
	{[]rune(".....+.58."), []Number{{58, 7, 8}}},
	{[]rune("..592....."), []Number{{592, 2, 4}}},
	{[]rune("......755."), []Number{{755, 6, 8}}},
	{[]rune("...$.*...."), []Number{}},
	{[]rune(".664.598.."), []Number{{664, 1, 3}, {598, 5, 7}}},
}

// var symbolTestCases = []struct {
// 	input          string
// 	expectedResult []Symbol
// }{
// 	{"467..114..", []Symbol{}},
// 	{"...*......", []Symbol{{"*", 3}}},
// 	{"..35..633.", []Symbol{}},
// 	{"......#...", []Symbol{{"#", 6}}},
// 	{"617*......", []Symbol{{"*", 3}}},
// 	{".....+.58.", []Symbol{{"+", 5}}},
// 	{"..592.....", []Symbol{}},
// 	{"......755.", []Symbol{}},
// 	{"...$.*....", []Symbol{{"$", 3}, {"*", 5}}},
// 	{".664.598..", []Symbol{}},
// }

// func TestCalculateSum(t *testing.T) {
// 	for _, testCase := range testCases {
// 		actualResult := calculateSum(testCase.input)
// 		if actualResult != testCase.expectedResult {
// 			t.Errorf("Actual result - %d is different from the expected the result - %d", actualResult, testCase.expectedResult)
// 		}

// 	}

// }

func TestFindNumber(t *testing.T) {
	for _, testCase := range numberTestCases {
		actualResult := findNumbers(testCase.input)

		for i, actual := range actualResult {
			if actual != testCase.expectedResult[i] {
				t.Errorf("Actual result - %v is different from the expected the result - %v", actual, testCase.expectedResult[i])
			}
		}

	}
}

// func TestFindSymbols(t *testing.T) {
// 	for _, testCase := range symbolTestCases {
// 		fmt.Println(testCase)
// 		actualResult := findSymbols(testCase.input)

// 		for i, actual := range actualResult {
// 			if actual != testCase.expectedResult[i] {
// 				t.Errorf("Actual result - %v is different from the expected the result - %v", actual, testCase.expectedResult[i])
// 			}
// 		}
// 	}
// }
