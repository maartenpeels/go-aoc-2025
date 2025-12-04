package day04

import (
	"strconv"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
}

func (s *Solution) Part1(input string) (string, error) {
	grid := utils.Grid(input)
	return strconv.Itoa(s.step(grid)), nil
}

func (s *Solution) Part2(input string) (string, error) {
	grid := utils.Grid(input)
	totalCount := 0

	for {
		count := s.step(grid)
		if count == 0 {
			break
		}
		totalCount += count

		spotsToRemove := [][2]int{}
		for x := 0; x < len(grid); x++ {
			for y := 0; y < len(grid[0]); y++ {
				if grid[x][y] == '@' {
					if s.neighbourCount(grid, x, y, '@') < 4 {
						spotsToRemove = append(spotsToRemove, [2]int{x, y})
					}
				}
			}
		}

		for _, spot := range spotsToRemove {
			x, y := spot[0], spot[1]
			grid[x][y] = 'x'
		}
	}

	return strconv.Itoa(totalCount), nil
}

func (s *Solution) step(grid [][]rune) int {
	totalCount := 0

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == '@' {
				count := s.neighbourCount(grid, x, y, '@')
				if count < 4 {
					totalCount++
				}
			}
		}
	}
	return totalCount
}

func (s *Solution) neighbourCount(grid [][]rune, x, y int, target rune) int {
	directions := [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	count := 0
	for _, dir := range directions {
		nx, ny := x+dir[0], y+dir[1]
		if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
			if grid[nx][ny] == target {
				count++
			}
		}
	}

	return count
}
