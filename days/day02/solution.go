package day02

import (
	"strconv"
	"strings"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
}

func (s *Solution) Part1(input string) (string, error) {
	return s.solve(input, true)
}

func (s *Solution) Part2(input string) (string, error) {
	return s.solve(input, false)
}

func (s *Solution) solve(input string, exactTwo bool) (string, error) {
	ranges := utils.SingleLineSplit(input, ",")
	invalidIDs := []int{}

	for _, r := range ranges {
		parts := strings.Split(r, "-")

		lower := strings.TrimSpace(parts[0])
		higher := strings.TrimSpace(parts[1])

		lowerInt, _ := strconv.Atoi(lower)
		higherInt, _ := strconv.Atoi(higher)

		for i := lowerInt; i <= higherInt; i++ {
			if s.isInvalid(i, exactTwo) {
				invalidIDs = append(invalidIDs, i)
			}
		}
	}

	total := 0
	for _, id := range invalidIDs {
		total += id
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) isInvalid(id int, exactTwo bool) bool {
	idStr := strconv.Itoa(id)
	idLen := len(idStr)

	for prefixLen := 1; prefixLen <= idLen/2; prefixLen++ {
		if idLen%prefixLen != 0 {
			continue
		}

		prefix := idStr[:prefixLen]
		repetitions := idLen / prefixLen

		if exactTwo && repetitions != 2 {
			continue
		}
		if !exactTwo && repetitions < 2 {
			continue
		}

		isValid := true
		for i := 0; i < repetitions; i++ {
			start := i * prefixLen
			end := start + prefixLen
			if idStr[start:end] != prefix {
				isValid = false
				break
			}
		}

		if isValid {
			return true
		}
	}

	return false
}
