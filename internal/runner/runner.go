package runner

import (
	"fmt"

	"github.com/maartenpeels/aoc-2025/days/day01"
	"github.com/maartenpeels/aoc-2025/days/day02"
	"github.com/maartenpeels/aoc-2025/days/day03"
	"github.com/maartenpeels/aoc-2025/internal/solver"
)

var solvers = map[string]solver.Solver{
	"01": &day01.Solution{},
	"02": &day02.Solution{},
	"03": &day03.Solution{},
}

func Run(day string, input string) error {
	s, ok := solvers[day]
	if !ok {
		return fmt.Errorf("no solution found for day %s", day)
	}

	fmt.Printf("=== Day %s ===\n\n", day)

	part1, err := s.Part1(input)
	if err != nil {
		return fmt.Errorf("part 1 error: %w", err)
	}
	fmt.Printf("Part 1: %s\n", part1)

	part2, err := s.Part2(input)
	if err != nil {
		return fmt.Errorf("part 2 error: %w", err)
	}
	fmt.Printf("Part 2: %s\n", part2)

	return nil
}
