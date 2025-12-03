package day03

import (
	"strconv"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
}

func (s *Solution) Part1(input string) (string, error) {
	banks := utils.LinesDigits(input)
	total := 0

	for _, bank := range banks {
		maxJoltage := 0

		for i := 0; i < len(bank); i++ {
			for j := i + 1; j < len(bank); j++ {
				joltage := bank[i]*10 + bank[j]
				if joltage > maxJoltage {
					maxJoltage = joltage
				}
			}
		}

		total += maxJoltage
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	banks := utils.LinesDigits(input)
	total := 0

	for _, bank := range banks {
		result := make([]int, 12)
		pos := 0

		for i := 0; i < 12; i++ {
			maxDigit := -1
			maxIdx := -1

			// Find the largest digit in the range that leaves enough digits for the remaining positions
			for j := pos; j <= len(bank)-(12-i); j++ {
				if bank[j] > maxDigit {
					maxDigit = bank[j]
					maxIdx = j
				}
			}

			result[i] = maxDigit
			pos = maxIdx + 1
		}

		joltage := 0
		for _, digit := range result {
			joltage = joltage*10 + digit
		}

		total += joltage
	}

	return strconv.Itoa(total), nil
}
