package utils

import "strings"

// Lines splits the input into an array of lines, trimming trailing newlines
func Lines(input string) []string {
	input = strings.TrimRight(input, "\n")
	if input == "" {
		return []string{}
	}
	return strings.Split(input, "\n")
}

// SingleLineSplit splits a single line input into an array based on the given separator
func SingleLineSplit(input string, sep string) []string {
	input = strings.TrimRight(input, "\n")
	return strings.Split(input, sep)
}
