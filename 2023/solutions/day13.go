package solutions

import (
	"fmt"
	"time"

  "aoc2023/utils"
)

func Day13() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day13")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

  part1, part2 := mirrors(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 13 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

func mirrors(lines []string) (int, int) {
  var rows []string
  sum1 := 0
  sum2 := 0

  for _, line := range lines {
    if len(line) > 0 {
      rows = append(rows, line)
    } else {
      cols := transposePattern(rows)

      sum1 += 100*reflection(rows) + reflection(cols)
      sum2 += 100*smudge(rows) + smudge(cols)

      rows = []string{}
    }
  }

  return sum1, sum2
}

func transposePattern (rows []string) []string {
  var cols []string

  for col := 0; col < len(rows[0]); col++ {
    line := ""
    for row := 0; row < len(rows); row++ {
      line += string(rows[row][col])
    }
    cols = append(cols, line)
  }

  return cols
}

func reflection(p []string) int {
  for i := 1; i < len(p); i++ {
    if checkReflection(p, i) {
      return i
    }
  }
  return 0
}

func checkReflection(p []string, start int) bool {
  isEqual := true
  var i, j = start, start - 1

  for {
    if p[i] != p[j] {
      isEqual = false
      break
    } 

    i++
    j--

    if i >= len(p) || j < 0 {
      break
    }
  }

  return isEqual
}

func smudge(p []string) int {
  for i := 1; i < len(p); i++ {
    if checkSmudge(p, i) {
      return i
    }
  }
  return 0
}

func checkSmudge(p []string, start int) bool {
  var i, j = start, start - 1
  diff := 0

  for {
    for col := 0; col < len(p[0]); col++ {
      if p[i][col] != p[j][col] {
        diff++
      }
    }

    i++
    j--

    if i >= len(p) || j < 0 || diff > 1 {
      break
    }
  }

  return diff == 1
}
