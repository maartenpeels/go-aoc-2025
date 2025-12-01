package day01

import (
	"fmt"
	"strconv"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
	dial        int
	zeroes      int
	part2Zeroes int
}

func (s *Solution) Part1(input string) (string, error) {
	lines := utils.Lines(input)
	s.zeroes = 0
	s.dial = 50

	// log.Println("The dial starts by pointing at " + strconv.Itoa(s.dial))

	for _, line := range lines {
		s.rotateDial(line)
	}

	return strconv.Itoa(s.zeroes), nil
}

func (s *Solution) Part2(input string) (string, error) {
	lines := utils.Lines(input)
	s.part2Zeroes = 0
	s.dial = 50

	for _, line := range lines {
		s.rotateDial(line)
	}

	return strconv.Itoa(s.part2Zeroes), nil
}

func (s *Solution) rotateDial(instruction string) {
	var direction rune = rune(instruction[0])
	var steps int
	fmt.Sscanf(instruction[1:], "%d", &steps)

	step := 1
	if direction == 'L' {
		step = -1
	}

	for i := 0; i < steps; i++ {
		s.dial = ((s.dial+step)%100 + 100) % 100
		if s.dial == 0 {
			s.part2Zeroes++
		}
	}

	if s.dial == 0 {
		s.zeroes++
	}

	// log.Println("The dial is rotated " + instruction + " to point at " + strconv.Itoa(s.dial))
}
