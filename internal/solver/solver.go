package solver

// Solver defines the interface that each day's solution must implement.
type Solver interface {
	Part1(input string) (string, error)
	Part2(input string) (string, error)
}
