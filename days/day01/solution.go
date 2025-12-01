package day01

import (
	"fmt"
	"log"
	"strconv"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
	dial   int
	zeroes int
}

func (s *Solution) Part1(input string) (string, error) {
	lines := utils.Lines(input)
	s.zeroes = 0
	s.dial = 50

	log.Println("The dial starts by pointing at " + strconv.Itoa(s.dial))

	for _, line := range lines {
		s.rotateDial(line)
		if s.dial == 0 {
			s.zeroes++
		}
	}

	return strconv.Itoa(s.zeroes), nil
}

func (s *Solution) Part2(input string) (string, error) {
	lines := utils.Lines(input)

	// TODO: Implement Part 2 solution
	_ = lines
	return "Not implemented yet", nil
}

func (s *Solution) rotateDial(instruction string) {
	var direction rune = rune(instruction[0])
	var steps int
	fmt.Sscanf(instruction[1:], "%d", &steps)

	if direction == 'L' {
		steps = -steps
	}

	s.dial = ((s.dial+steps)%100 + 100) % 100

	log.Println("The dial is rotated " + instruction + " to point at " + strconv.Itoa(s.dial))
}
