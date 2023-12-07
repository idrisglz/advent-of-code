package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	text     string
	expected bool
}{
	{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green", true},
	{"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue", true},
	{"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red", false},
	{"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red", false},
	{"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green", true},
}

func TestIsEligible(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Checking %s", testCase.text), func(t *testing.T) {
			actual_result := isEligible(testCase.text)
			if actual_result != testCase.expected {
				t.Errorf("Result is not equal to the expected result, %s - %t vs %t", testCase.text, actual_result, testCase.expected)
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
