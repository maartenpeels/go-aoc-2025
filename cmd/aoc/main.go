package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/maartenpeels/aoc-2025/internal/runner"
	"github.com/spf13/cobra"
)

var (
	useTest bool
)

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code 2025 solutions",
	Long:  `A CLI tool to run Advent of Code 2025 solutions for any day.`,
}

var dayCmd = &cobra.Command{
	Use:   "day [day_number]",
	Short: "Run solution for a specific day",
	Long:  `Run the Advent of Code solution for a specific day (1-25).`,
	Args:  cobra.ExactArgs(1),
	Run:   runDay,
}

func init() {
	dayCmd.Flags().BoolVar(&useTest, "test", false, "Use example.txt instead of input.txt")
	rootCmd.AddCommand(dayCmd)
}

func runDay(cmd *cobra.Command, args []string) {
	day := args[0]

	// Pad single-digit days with leading zero
	if len(day) == 1 {
		day = "0" + day
	}

	inputFile := "input.txt"
	if useTest {
		inputFile = "example.txt"
	}

	inputPath := filepath.Join("days", fmt.Sprintf("day%s", day), inputFile)

	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		os.Exit(1)
	}

	input := string(data)

	if err := runner.Run(day, input); err != nil {
		fmt.Fprintf(os.Stderr, "Error running day %s: %v\n", day, err)
		os.Exit(1)
	}
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
