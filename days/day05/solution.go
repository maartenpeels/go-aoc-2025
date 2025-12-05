package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/maartenpeels/aoc-2025/internal/utils"
)

type Solution struct {
}

type Range struct {
	Start int
	End   int
}

func (s *Solution) Part1(input string) (string, error) {
	freshIngredientRanges, availableIngredients := parseInput(input)
	freshCount := 0

	for _, ingredient := range availableIngredients {
		isFresh := false
		for _, r := range freshIngredientRanges {
			if ingredient >= r.Start && ingredient <= r.End {
				isFresh = true
				break
			}
		}
		if isFresh {
			freshCount++
		}
	}

	return strconv.Itoa(freshCount), nil
}

func (s *Solution) Part2(input string) (string, error) {
	freshIngredientRanges, _ := parseInput(input)
	freshIngredients := 0

	mergedRanges := mergeRanges(freshIngredientRanges)
	for _, r := range mergedRanges {
		freshIngredients += (r.End - r.Start + 1)
	}

	return strconv.Itoa(freshIngredients), nil
}

func parseInput(input string) ([]Range, []int) {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")
	freshIngedientRangeLines := utils.Lines(parts[0])
	availableIngredientLines := utils.Lines(parts[1])

	freshIngredientRanges := make([]Range, len(freshIngedientRangeLines))
	availableIngredients := make([]int, len(availableIngredientLines))

	for i, line := range freshIngedientRangeLines {
		rangeParts := strings.Split(line, "-")
		start := utils.Atoi(rangeParts[0])
		end := utils.Atoi(rangeParts[1])
		freshIngredientRanges[i] = Range{Start: start, End: end}
	}

	for i, line := range availableIngredientLines {
		availableIngredients[i] = utils.Atoi(line)
	}

	return freshIngredientRanges, availableIngredients
}

func mergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	merged := []Range{}
	merged = append(merged, ranges[0])

	for i := 1; i < len(ranges); i++ {
		lastMerged := &merged[len(merged)-1]
		current := ranges[i]

		if current.Start <= lastMerged.End {
			if current.End > lastMerged.End {
				lastMerged.End = current.End
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}
