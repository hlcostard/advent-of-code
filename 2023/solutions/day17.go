package solutions

import (
	"fmt"
	"time"

	"aoc2023/utils"
)

func Day17() {
	start := time.Now()

	lines, err := utils.ReadRunes("test-input/day17")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

  nums := utils.RunesToInt(lines)

  part1 := moveCrucible(nums)
  part2 := moveCrucible(nums)

	elapsed := time.Since(start)
	fmt.Printf("Day 17 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type block struct {
  i         int
  j         int
  visited   bool
  shortDist int
  prevBlock string
}

func moveCrucible(lines [][]int) int {





  return 420
}
