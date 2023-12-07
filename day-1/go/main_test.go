package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	text     string
	expected int
}{
	{"1abc2", 12},
	{"pqr3stu8vwx", 38},
	{"a1b2c3d4e5f", 15},
	{"treb7uchet", 77},
	{"two1nine", 29},
	{"eightwothree", 83},
	{"abcone2threexyz", 13},
	{"xtwone3four", 24},
	{"4nineeightseven2", 42},
	{"zoneight234", 14},
	{"7pqrstsixteen", 76},
	{"1abc2", 12},
	{"pqr3stu8vwx", 38},
	{"a1b2c3d4e5f", 15},
	{"treb7uchet", 77},
	{"eightsixone1qzp4nxzrslmzrsix", 86},
	{"hsconethree1fourthree", 13},
	{"32eightseventhreekqdgtcxgjxvv1seven", 37},
	{"eighthree", 83},
	{"eightwo", 82},
}

func TestGetCalibrationValues(t *testing.T) {

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Checking %s", testCase.text), func(t *testing.T) {
			actual_result, _ := getCalibrationValues(testCase.text)
			if actual_result != testCase.expected {
				t.Errorf("Result is not equal to the expected result, %s - %d vs %d", testCase.text, actual_result, testCase.expected)
			}

		})
	}

}

func BenchmarkGetCalibrationValues(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			getCalibrationValues(testCase.text)
		}
	}
}

// Run with: go test -bench=BenchmarkMyFunction
