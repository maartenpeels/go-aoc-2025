package utils

import (
	"strings"
)

// Lines splits the input into an array of lines, trimming trailing newlines
func Lines(input string) []string {
	input = strings.TrimRight(input, "\n")
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "\n")
}

// LinesDigits splits the input into lines, and each line into individual digits
func LinesDigits(input string) [][]int {
	lines := Lines(input)
	result := make([][]int, len(lines))
	for i, line := range lines {
		digits := make([]int, len(line))
		for j, ch := range line {
			digits[j] = int(ch - '0')
		}
		result[i] = digits
	}
	return result
}

// SingleLineSplit splits a single line input into an array based on the given separator
func SingleLineSplit(input string, sep string) []string {
	input = strings.TrimRight(input, "\n")
	return strings.Split(input, sep)
}
