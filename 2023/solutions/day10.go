package solutions

import (
	"fmt"
	"slices"
	"time"

	"aoc2023/utils"
)

func Day10() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day10")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1, part2 := stepsAndAreaLoop(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 10 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type path struct {
	line int
	col  int
	pipe string
}

func stepsAndAreaLoop(input []string) (int, int) {
	var paths []path
	looped := false
	count := 0

	line, col, pipe, direction, positionS := findS(input)

	paths = append(paths, positionS)

	for !looped {
		count++

		paths = append(paths, path{line: line, col: col, pipe: pipe})
		pipe, direction, line, col = nextPipe(input, direction, pipe, line, col)

		if input[line][col] == 'S' {
			looped = true
		}
	}

	countArea := 0
	verticalPipes := []string{"|", "7", "F", "J", "L"}

	for i := 0; i < len(input)-1; i++ {
		lastVertPipe := ""
		isRegion := false
		countMain := 0

		for j := 0; j < len(input[i]); j++ {
			if isMainLoop(i, j, paths) {
				countMain++
				p := string(input[i][j])
				if p == "S" {
					p = positionS.pipe
				}

				if slices.Contains(verticalPipes, p) {
					if p == "|" {
						isRegion = !isRegion
					} else if p == "J" && lastVertPipe == "F" {
						isRegion = !isRegion
					} else if p == "7" && lastVertPipe == "L" {
						isRegion = !isRegion
					}
					lastVertPipe = p
				}
			} else {
				if isRegion {
					countArea++
				}
			}
		}
	}

	return (count + 1) / 2, countArea
}

func isMainLoop(i, j int, paths []path) bool {
	for _, p := range paths {
		if p.line == i && p.col == j {
			return true
		}
	}

	return false
}

func findS(input []string) (int, int, string, string, path) {
	var line int
	var col int
	var lineS int
	var colS int
	var pipe []string
	var direction []string

	for i, lineInput := range input {
		for j := 0; j < len(lineInput); j++ {
			if lineInput[j] == 'S' {
				lineS = i
				colS = j
				line = i
				col = j
				if i > 0 {
					north := input[i-1][j]
					if north == '|' || north == '7' || north == 'F' {
						pipe = append(pipe, string(north))
						direction = append(direction, "south")
						line--
					}
				}
				if i < len(input)-1 {
					south := input[i+1][j]
					if south == '|' || south == 'J' || south == 'L' {
						pipe = append(pipe, string(south))
						direction = append(direction, "north")
						line++
					}
				}
				if j > 0 {
					west := input[i][j-1]
					if west == '-' || west == 'F' || west == 'L' {
						pipe = append(pipe, string(west))
						direction = append(direction, "east")
					}
				}
				if j < len(lineInput)-1 {
					east := input[i][j+1]
					if east == '-' || east == 'J' || east == '7' {
						pipe = append(pipe, string(east))
						direction = append(direction, "west")
					}
				}
			}
		}
	}

	var pipeS string
	if direction[0] == "south" && direction[1] == "east" {
		pipeS = "L"
	} else if direction[0] == "south" && direction[1] == "west" {
		pipeS = "J"
	} else if direction[0] == "north" && direction[1] == "east" {
		pipeS = "7"
	} else if direction[0] == "north" && direction[1] == "west" {
		pipeS = "F"
	}

	positionS := path{line: lineS, col: colS, pipe: pipeS}
	return line, col, pipe[0], direction[0], positionS
}

func nextPipe(input []string, direction, pipe string, i, j int) (string, string, int, int) {
	var nextPipe string
	var nextDirection string
	nextLine := i
	nextCol := j

	if direction == "north" {
		if pipe == "|" {
			nextLine = i + 1
			nextDirection = "north"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "J" {
			nextCol = j - 1
			nextDirection = "east"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "L" {
			nextCol = j + 1
			nextDirection = "west"
			nextPipe = string(input[nextLine][nextCol])
		}
		return nextPipe, nextDirection, nextLine, nextCol
	}

	if direction == "west" {
		if pipe == "-" {
			nextCol = j + 1
			nextDirection = "west"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "J" {
			nextLine = i - 1
			nextDirection = "south"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "7" {
			nextLine = i + 1
			nextDirection = "north"
			nextPipe = string(input[nextLine][nextCol])
		}
		return nextPipe, nextDirection, nextLine, nextCol
	}

	if direction == "east" {
		if pipe == "-" {
			nextCol = j - 1
			nextDirection = "east"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "L" {
			nextLine = i - 1
			nextDirection = "south"
			nextPipe = string(input[nextLine][nextCol])
		} else if pipe == "F" {
			nextLine = i + 1
			nextDirection = "north"
			nextPipe = string(input[nextLine][nextCol])
		}
		return nextPipe, nextDirection, nextLine, nextCol
	}

	// direction == south
	if pipe == "|" {
		nextLine = i - 1
		nextDirection = "south"
		nextPipe = string(input[nextLine][nextCol])
	} else if pipe == "7" {
		nextCol = j - 1
		nextDirection = "east"
		nextPipe = string(input[nextLine][nextCol])
	} else if pipe == "F" {
		nextCol = j + 1
		nextDirection = "west"
		nextPipe = string(input[nextLine][nextCol])
	}

	return nextPipe, nextDirection, nextLine, nextCol
}
