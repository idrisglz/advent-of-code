package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	text           string
	expectedResult int
}{
	{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", 48},
	{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", 12},
	{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", 1560},
	{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", 630},
	{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", 36},
}

func TestIsEligible(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Checking %s", testCase.text), func(t *testing.T) {
			green, blue, red := getMaxCount(testCase.text, greenRegex, blueRegex, redRegex)
			actualResult := green * blue * red
			if actualResult != testCase.expectedResult {
				t.Errorf("Result is not equal to the expected result, %s - %d vs %d", testCase.text, actualResult, testCase.expectedResult)
			}
		})
	}

}

// func BenchmarkGetCalibrationValues(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, testCase := range testCases {
// 			getCalibrationValues(testCase.text)
// 		}
// 	}
// }
