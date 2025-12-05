package utils

import "strconv"

// Atoi without error handling
func Atoi(s string) int {
	result, _ := strconv.Atoi(s)
	return result
}
