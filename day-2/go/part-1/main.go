package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var inputs []string = []string{
	"Game 1: 6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blue",
	"Game 2: 2 red, 7 green; 13 green, 2 blue, 4 red; 4 green, 5 red, 1 blue; 1 blue, 9 red, 1 green",
	"Game 3: 2 green, 3 blue, 9 red; 3 red, 2 green; 6 red, 4 blue; 6 red",
	"Game 4: 9 red, 3 green; 3 green, 8 red, 6 blue; 12 blue, 4 green, 6 red; 4 green, 18 blue, 11 red; 9 blue, 2 green, 3 red; 14 blue, 7 red",
	"Game 5: 1 blue, 2 green, 3 red; 16 red, 6 green; 6 green, 2 red; 9 red, 1 green",
	"Game 6: 4 green, 7 red, 1 blue; 18 green, 6 blue, 7 red; 1 blue, 3 red, 9 green; 9 red, 19 green, 1 blue; 7 red, 9 green, 4 blue; 5 red, 5 blue, 10 green",
	"Game 7: 16 blue, 5 green, 6 red; 1 blue, 6 red, 9 green; 6 green, 3 red, 2 blue; 2 red, 12 blue, 2 green",
	"Game 8: 6 green, 10 red; 7 red, 6 green, 17 blue; 13 blue, 1 red",
	"Game 9: 2 red, 4 green, 5 blue; 2 green, 5 blue; 4 green, 1 blue; 3 green, 3 red, 6 blue; 3 green",
	"Game 10: 3 green, 5 red, 6 blue; 4 blue, 4 red, 5 green; 5 green, 9 red, 5 blue; 4 green, 6 blue, 10 red",
	"Game 11: 1 blue, 7 red, 9 green; 1 blue, 13 red, 7 green; 5 red; 4 red, 7 green, 2 blue; 7 green, 12 red; 13 red, 2 blue, 12 green",
	"Game 12: 4 blue, 2 red; 9 blue, 2 green, 3 red; 8 blue, 1 green, 1 red; 2 red, 3 green, 11 blue",
	"Game 13: 6 red, 8 green, 2 blue; 6 red, 7 green; 3 green, 3 red; 2 blue; 3 red, 5 green",
	"Game 14: 2 green, 11 blue, 1 red; 5 blue, 1 red, 1 green; 3 green, 12 blue, 2 red",
	"Game 15: 4 blue, 6 red, 7 green; 1 red, 2 blue, 5 green; 6 red, 3 green, 8 blue; 7 green, 8 blue, 4 red",
	"Game 16: 2 red, 16 blue; 2 green, 7 red; 15 blue, 7 red; 2 red, 3 green, 3 blue; 3 red, 1 green, 4 blue; 4 blue, 3 green",
	"Game 17: 2 red, 3 green, 10 blue; 9 red, 4 blue, 3 green; 3 green, 11 red, 6 blue",
	"Game 18: 1 red; 6 green, 1 red, 9 blue; 4 blue, 2 green; 6 blue, 10 green",
	"Game 19: 2 red, 9 green; 2 red, 1 blue; 5 blue, 12 green; 5 green; 8 green, 2 red, 3 blue; 1 red, 11 green",
	"Game 20: 3 green, 2 red, 7 blue; 1 blue, 10 green; 1 red, 14 blue, 13 green; 3 green, 19 blue, 4 red",
	"Game 21: 8 red, 10 blue, 8 green; 2 red, 7 green, 18 blue; 4 green, 11 blue, 4 red; 5 green, 3 blue, 10 red",
	"Game 22: 17 blue, 2 green, 2 red; 8 red, 7 blue; 1 red, 9 blue, 1 green",
	"Game 23: 4 blue, 18 red, 4 green; 3 blue, 7 red; 11 red; 3 blue, 6 red; 19 red",
	"Game 24: 10 red, 1 blue, 17 green; 17 green, 6 red; 14 green, 4 blue",
	"Game 25: 4 blue, 9 green, 4 red; 3 green, 5 blue; 5 blue, 8 green; 3 green, 3 blue, 1 red; 10 green, 1 blue, 4 red; 2 green, 2 blue, 1 red",
	"Game 26: 18 green, 3 red, 12 blue; 2 red, 7 green; 11 blue, 17 green; 12 green, 11 blue; 12 green, 4 blue, 3 red",
	"Game 27: 1 red, 3 blue, 8 green; 15 blue, 8 red, 4 green; 6 red, 9 blue; 6 red, 12 blue, 9 green; 4 red, 7 blue, 15 green",
	"Game 28: 1 red, 14 green; 1 blue, 11 green; 2 green; 4 red, 6 green, 1 blue",
	"Game 29: 1 green, 13 red; 4 red, 16 green, 7 blue; 2 red, 4 blue; 12 green, 8 blue, 4 red; 2 red; 12 red, 5 green",
	"Game 30: 3 green, 4 blue, 3 red; 5 blue, 4 green, 7 red; 5 blue, 2 green, 2 red; 3 red, 1 blue",
	"Game 31: 1 blue, 8 green; 9 green, 2 blue, 1 red; 1 red, 2 blue",
	"Game 32: 11 red, 5 green, 4 blue; 3 blue, 11 red, 8 green; 6 blue, 3 green, 17 red; 4 red, 7 green, 10 blue",
	"Game 33: 6 blue, 4 red; 1 green; 1 green, 4 red, 4 blue; 1 green, 3 red, 10 blue; 10 blue, 1 red",
	"Game 34: 2 green, 3 blue, 3 red; 4 red; 2 red, 2 blue",
	"Game 35: 9 green, 13 blue; 13 blue, 14 red, 1 green; 11 blue, 4 red, 7 green; 5 blue, 5 red, 8 green; 4 red, 2 blue, 2 green",
	"Game 36: 9 red, 5 blue, 8 green; 7 red, 20 blue; 6 green, 16 blue, 5 red; 12 red, 3 blue, 3 green; 3 green, 6 blue, 11 red; 11 red, 8 blue, 3 green",
	"Game 37: 10 green, 11 red, 3 blue; 2 blue, 6 green, 11 red; 9 green, 8 red, 2 blue",
	"Game 38: 2 red, 2 green, 4 blue; 3 red, 4 green, 3 blue; 8 green, 1 blue, 1 red; 3 red, 5 blue, 5 green",
	"Game 39: 3 green, 17 red, 4 blue; 2 green, 20 red; 4 blue, 4 red, 5 green; 5 blue, 7 green, 7 red; 4 blue, 5 green, 16 red",
	"Game 40: 2 green, 2 blue, 4 red; 3 blue, 16 green; 1 green, 2 blue; 1 red; 3 blue, 15 green; 13 green, 1 red, 2 blue",
	"Game 41: 12 red, 10 blue, 9 green; 1 green, 15 red, 4 blue; 2 green, 8 blue, 12 red; 3 red, 4 green, 2 blue; 8 blue, 14 red, 10 green; 9 blue, 7 green, 6 red",
	"Game 42: 5 red, 3 green, 6 blue; 4 blue, 6 green, 2 red; 10 blue; 3 red, 6 green, 10 blue",
	"Game 43: 9 blue, 7 green, 1 red; 2 green, 2 red, 8 blue; 3 red, 15 blue, 11 green; 1 red, 6 blue, 1 green; 2 red, 1 blue; 1 red, 3 green, 7 blue",
	"Game 44: 4 green, 6 red; 15 green, 6 red; 9 green, 16 red, 7 blue; 11 green, 4 blue, 12 red",
	"Game 45: 3 blue, 6 green, 1 red; 4 green, 3 blue; 8 green, 3 blue",
	"Game 46: 10 red, 8 blue; 12 red, 2 green, 17 blue; 17 blue, 6 red, 1 green; 18 red, 6 green, 3 blue; 16 blue, 2 green, 3 red",
	"Game 47: 8 green, 13 red; 8 green, 8 red, 4 blue; 10 red, 3 green; 14 red, 5 green, 8 blue; 7 green, 19 red, 3 blue; 2 red, 5 green, 5 blue",
	"Game 48: 7 green, 9 blue, 3 red; 7 blue, 1 green, 9 red; 7 green, 4 red, 1 blue; 6 green, 3 red, 1 blue",
	"Game 49: 2 red, 3 green; 3 blue, 2 red; 4 red, 3 blue",
	"Game 50: 3 red, 7 blue, 4 green; 2 green, 1 blue, 7 red; 4 red, 1 green, 5 blue",
	"Game 51: 11 red, 6 green, 1 blue; 7 red, 1 blue, 9 green; 15 red, 18 green; 11 green, 1 blue, 11 red; 10 green, 14 red; 1 red, 11 green, 1 blue",
	"Game 52: 18 blue, 1 red, 2 green; 18 blue, 3 green, 1 red; 2 green, 13 blue, 1 red",
	"Game 53: 2 blue, 9 red, 6 green; 1 blue, 3 red; 7 red, 6 blue, 8 green; 2 red, 3 blue, 4 green; 1 green, 2 blue, 2 red",
	"Game 54: 16 red, 4 blue; 1 green, 3 blue, 3 red; 2 green, 12 red; 2 green, 1 blue, 3 red; 10 blue, 6 red, 2 green",
	"Game 55: 1 blue, 4 red, 1 green; 2 blue, 2 red; 13 red, 4 blue, 1 green; 4 blue, 9 red; 1 green, 1 blue, 16 red",
	"Game 56: 12 blue, 12 green; 4 blue, 1 red, 3 green; 2 red, 12 green; 1 red, 11 green, 13 blue; 16 blue, 5 green",
	"Game 57: 1 blue, 3 red; 10 green, 5 red; 5 green, 2 red; 1 red, 13 green",
	"Game 58: 6 blue, 1 red, 6 green; 3 red, 9 blue; 4 red, 9 blue, 5 green; 1 green, 5 red, 7 blue",
	"Game 59: 10 red, 3 green, 3 blue; 6 blue, 11 red, 1 green; 5 green, 10 red; 16 red, 2 blue, 4 green; 3 green, 10 red",
	"Game 60: 2 green, 1 blue; 1 green, 1 blue, 4 red; 3 blue, 1 red, 1 green; 2 red, 2 green; 4 red",
	"Game 61: 5 red, 1 green, 10 blue; 9 red, 10 blue; 1 red, 2 green, 4 blue; 10 blue, 2 green, 9 red; 1 red, 12 blue, 2 green",
	"Game 62: 1 blue, 5 green; 4 blue, 12 green, 1 red; 7 blue, 3 green; 7 blue, 3 green; 3 blue, 1 green, 2 red; 7 blue, 1 red, 12 green",
	"Game 63: 4 blue, 2 green, 5 red; 1 green, 2 red, 2 blue; 4 blue, 2 red, 2 green; 1 blue, 6 red, 2 green; 6 blue, 1 red; 1 green, 9 red, 6 blue",
	"Game 64: 1 green; 3 green, 5 blue, 5 red; 3 red, 3 blue, 3 green; 1 green, 4 red, 6 blue; 5 red",
	"Game 65: 2 red; 1 blue, 1 red; 7 red, 2 blue; 1 green, 4 blue, 3 red",
	"Game 66: 3 red, 9 blue; 1 red, 6 blue, 15 green; 3 green, 3 red, 11 blue",
	"Game 67: 2 red, 1 green, 2 blue; 6 red, 1 green; 1 blue, 1 red, 4 green",
	"Game 68: 3 red, 1 blue; 1 green, 3 red, 2 blue; 1 green, 8 red; 2 blue, 3 red",
	"Game 69: 5 blue, 6 red; 1 green, 15 blue, 10 red; 1 green, 2 red, 4 blue; 5 blue, 7 red; 3 red, 1 green, 11 blue",
	"Game 70: 4 green, 2 red, 8 blue; 5 red, 3 blue; 10 green, 5 blue",
	"Game 71: 1 red, 2 blue, 9 green; 3 red, 8 green; 1 red, 2 blue, 6 green; 3 red, 6 blue, 8 green; 6 green, 3 blue, 2 red; 3 red, 8 green, 6 blue",
	"Game 72: 1 red, 11 green; 1 blue, 7 green, 1 red; 2 red, 12 green; 10 green, 6 red",
	"Game 73: 9 green, 2 red; 1 blue, 3 green; 1 blue, 1 red, 7 green; 2 blue, 9 green, 4 red; 2 blue, 3 red, 8 green; 2 green, 9 red",
	"Game 74: 2 green, 7 red; 1 green, 3 blue, 6 red; 4 green, 3 blue, 6 red; 2 green, 3 blue, 1 red; 3 red, 2 blue, 1 green",
	"Game 75: 15 green, 2 blue; 15 green, 6 red, 2 blue; 12 green, 2 blue, 1 red",
	"Game 76: 1 red, 9 green, 12 blue; 6 red, 12 green, 1 blue; 7 green, 2 blue, 1 red",
	"Game 77: 11 blue, 1 red; 7 blue, 2 red, 13 green; 10 blue, 10 green; 12 blue, 3 red",
	"Game 78: 4 green; 1 blue, 5 green; 5 green, 1 blue, 1 red",
	"Game 79: 4 green, 7 blue, 16 red; 1 blue, 10 red, 5 green; 3 green, 4 red, 3 blue; 11 blue, 18 red, 5 green",
	"Game 80: 1 red, 4 blue, 6 green; 14 blue, 16 red, 2 green; 2 blue, 5 red, 4 green; 2 green, 8 red; 18 red, 6 green, 2 blue; 18 red, 9 blue",
	"Game 81: 11 red, 8 blue, 1 green; 12 blue, 2 green, 14 red; 16 red, 2 green, 6 blue; 17 red, 2 green; 3 green, 3 blue, 15 red",
	"Game 82: 13 red, 1 blue, 6 green; 3 green, 12 red, 3 blue; 5 red, 3 green, 18 blue; 15 blue, 8 red",
	"Game 83: 9 green, 5 blue, 5 red; 8 green, 15 blue, 7 red; 4 green, 6 red, 10 blue; 5 green, 2 red",
	"Game 84: 2 blue, 2 green, 6 red; 2 green, 7 red, 1 blue; 3 green, 3 blue; 2 green, 3 red, 3 blue; 6 green, 4 red",
	"Game 85: 1 blue, 3 green, 5 red; 2 green, 2 red; 4 red, 3 blue; 2 green, 3 blue, 1 red; 4 red, 2 green, 4 blue",
	"Game 86: 6 red, 1 blue; 1 green, 16 red; 2 green, 1 red; 12 red, 1 blue",
	"Game 87: 6 red, 12 green, 1 blue; 5 blue, 6 red, 4 green; 2 blue, 5 red, 8 green",
	"Game 88: 3 green, 6 red, 2 blue; 3 blue, 2 green, 6 red; 1 red, 11 blue, 2 green",
	"Game 89: 7 red, 3 blue, 9 green; 6 red, 3 blue, 15 green; 2 blue, 6 red, 12 green; 5 red, 8 green; 3 blue, 7 red, 9 green; 5 red, 7 green",
	"Game 90: 2 green, 4 red, 19 blue; 13 blue, 4 red, 1 green; 14 blue, 8 green",
	"Game 91: 12 green, 5 blue, 4 red; 9 green, 10 blue, 1 red; 13 green, 1 blue, 4 red; 2 red, 5 blue; 2 blue, 7 green, 2 red; 5 blue, 5 green, 3 red",
	"Game 92: 9 red, 6 blue, 16 green; 11 green, 2 red, 7 blue; 1 green, 1 red, 3 blue; 4 green, 8 red",
	"Game 93: 1 green, 4 blue, 8 red; 2 red, 1 green, 2 blue; 2 blue, 9 red; 1 green, 4 blue, 3 red; 3 red, 1 green, 4 blue",
	"Game 94: 1 green, 7 red, 4 blue; 4 red, 3 blue; 16 blue, 9 red, 7 green; 9 red, 15 blue; 15 blue, 3 red, 6 green; 7 red, 10 blue, 12 green",
	"Game 95: 5 green, 6 blue; 10 green, 9 blue; 4 blue, 8 green, 2 red; 5 blue, 5 green, 1 red",
	"Game 96: 13 blue, 10 red, 2 green; 10 red, 2 green, 1 blue; 6 blue, 5 red, 3 green; 11 red, 3 green, 5 blue; 11 red, 2 green; 3 green, 6 blue",
	"Game 97: 9 green, 11 red, 8 blue; 6 red, 9 blue, 2 green; 3 red, 17 blue, 1 green",
	"Game 98: 14 blue, 3 green; 2 red, 15 blue, 3 green; 15 blue, 8 green, 1 red; 1 red, 8 green",
	"Game 99: 2 green, 7 blue; 14 red, 1 green, 4 blue; 8 blue, 13 red, 2 green; 10 green, 7 red, 10 blue",
	"Game 100: 5 green, 11 blue, 6 red; 5 green, 12 blue; 1 green, 14 blue, 1 red; 3 blue, 5 red, 6 green; 9 blue; 6 red",
}

func main() {

	eligibleIds := []int{}

	for _, input := range inputs {
		id, err := getId(input)
		if err != nil {
			os.Exit(1)
		}

		if isEligible(input) {
			eligibleIds = append(eligibleIds, id)
		}

	}

	sum := 0
	for _, id := range eligibleIds {
		sum += id
	}
	fmt.Println(sum)
}

func isEligible(input string) bool {
	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	startIndex := strings.Index(input, ":")
	values_input := input[startIndex+1:]
	parts := strings.Split(values_input, ";")

	for _, p := range parts {
		green, blue, red, err := countColors(p)
		if err != nil || green > greenLimit || red > redLimit || blue > blueLimit {
			return false
		}
	}
	return true
}

func getId(input string) (id int, err error) {
	idRegex := regexp.MustCompile(`Game\s*(\d+):`)

	idMatch := idRegex.FindStringSubmatch(input)
	if idMatch != nil {
		fmt.Println(idMatch)
		id, err = strconv.Atoi(idMatch[1])
		if err != nil {
			return 0, err
		}

	}
	return id, nil
}

func countColors(input string) (green int, blue int, red int, err error) {
	greenRegex := regexp.MustCompile(`(\d+)\s*green`)
	blueRegex := regexp.MustCompile(`(\d+)\s*blue`)
	redRegex := regexp.MustCompile(`(\d+)\s*red`)

	greenMatch := greenRegex.FindStringSubmatch(input)
	if greenMatch != nil {
		green, err = strconv.Atoi(greenMatch[1])
		if err != nil {
			return 0, 0, 0, err
		}
	}

	blueMatch := blueRegex.FindStringSubmatch(input)
	if blueMatch != nil {
		blue, err = strconv.Atoi(blueMatch[1])
		if err != nil {
			return 0, 0, 0, err
		}
	}

	redMatch := redRegex.FindStringSubmatch(input)
	if redMatch != nil {
		red, err = strconv.Atoi(redMatch[1])
		if err != nil {
			return 0, 0, 0, err
		}
	}

	return green, blue, red, nil
}
