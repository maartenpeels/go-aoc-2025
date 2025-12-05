package create

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func Run(day string) error {
	err := createDayFiles(day)
	if err != nil {
		return err
	}

	err = addSolverToRunner(day)
	if err != nil {
		return err
	}

	return nil
}

func createDayFiles(day string) error {
	pathBase := path.Join("days", "day"+day)
	exampleFilePath := path.Join(pathBase, "example.txt")
	inputFilePath := path.Join(pathBase, "input.txt")
	solutionFilePath := path.Join(pathBase, "solution.go")

	err := createFileIfNotExists(exampleFilePath, "")
	if err != nil {
		return err
	}

	err = createFileIfNotExists(inputFilePath, "")
	if err != nil {
		return err
	}

	err = createFileIfNotExists(solutionFilePath, generateSolutionFileContent(day))
	if err != nil {
		return err
	}

	return nil
}

func createFileIfNotExists(filePath string, content string) error {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filePath), os.ModePerm)
		if err != nil {
			return err
		}

		err = os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func generateSolutionFileContent(day string) string {
	return `package day` + day + `

type Solution struct {
}

func (s *Solution) Part1(input string) (string, error) {
	// Implement Part 1 solution here
	return "", nil
}

func (s *Solution) Part2(input string) (string, error) {
	// Implement Part 2 solution here
	return "", nil
}
`
}

func addSolverToRunner(day string) error {
	runnerFilePath := path.Join("internal", "runner", "runner.go")

	contentBytes, err := os.ReadFile(runnerFilePath)
	if err != nil {
		return err
	}
	content := string(contentBytes)

	dayInt, err := strconv.Atoi(day)
	if err != nil {
		return err
	}

	importLine := "\t\"github.com/maartenpeels/aoc-2025/days/day" + day + "\"\n"
	mapEntryLine := "\t\"" + day + "\": &day" + day + ".Solution{},\n"

	if dayInt == 1 {
		importInsertPoint := "import (\n"
		content = insertAfterLine(content, importInsertPoint, importLine)

		mapInsertPoint := "var solvers = map[string]solver.Solver{\n"
		content = insertAfterLine(content, mapInsertPoint, mapEntryLine)
	} else {
		prevDay := fmt.Sprintf("%02d", dayInt-1)
		importInsertPoint := "\t\"github.com/maartenpeels/aoc-2025/days/day" + prevDay + "\"\n"
		content = insertAfterLine(content, importInsertPoint, importLine)

		mapInsertPoint := "\t\"" + prevDay + "\": &day" + prevDay + ".Solution{},\n"
		content = insertAfterLine(content, mapInsertPoint, mapEntryLine)
	}

	// Write the updated content back to runner.go
	err = os.WriteFile(runnerFilePath, []byte(content), 0644)
	if err != nil {
		return err
	}

	return nil
}

func insertAfterLine(content string, lineToFind string, lineToInsert string) string {
	return strings.Replace(content, lineToFind, lineToFind+lineToInsert, 1)
}
