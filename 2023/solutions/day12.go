package solutions

import (
	"fmt"
	"time"

	"aoc2023/utils"
)

func Day12() {
	start := time.Now()

	lines, err := utils.ReadLines("test-input/day12")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1, part2 := shortestPathGalaxies(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 12 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}
