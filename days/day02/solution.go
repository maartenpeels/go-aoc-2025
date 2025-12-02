package day02

import (
	"log"
	"strconv"
	"strings"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
	invalidIDs []int
}

func (s *Solution) Part1(input string) (string, error) {
	ranges := utils.SingleLineSplit(input, ",")

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {
			log.Println(parts)
			log.Fatalln("Error: malformed input")
			continue
		}

		lower := strings.TrimSpace(parts[0])
		higher := strings.TrimSpace(parts[1])

		lowerInt, errLower := strconv.Atoi(lower)
		higherInt, errHigher := strconv.Atoi(higher)

		if errLower != nil || errHigher != nil {
			log.Println(lower, higher)
			log.Fatalln("Error: malformed input")
			continue
		}

		for i := lowerInt; i <= higherInt; i++ {
			if s.isInvalid(i) {
				s.invalidIDs = append(s.invalidIDs, i)
			}
		}
	}

	total := 0
	for _, id := range s.invalidIDs {
		total += id
	}

	return strconv.Itoa(total), nil
}

func (s *Solution) Part2(input string) (string, error) {
	//ranges := utils.SingleLineSplit(input, ",")

	return "", nil
}

func (s *Solution) isInvalid(id int) bool {
	// Needs to be even length (chars)
	if len(strconv.Itoa(id))%2 != 0 {
		return false
	}

	// First half of digits need to be the same as the second half
	firstHalf := strconv.Itoa(id)[:len(strconv.Itoa(id))/2]
	secondHalf := strconv.Itoa(id)[len(strconv.Itoa(id))/2:]

	return firstHalf == secondHalf
}
