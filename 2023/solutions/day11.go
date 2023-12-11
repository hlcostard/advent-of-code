package solutions

import (
	"fmt"
	"slices"
	"time"

	"aoc2023/utils"
)

func Day11() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day11")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1, part2 := shortestPathGalaxies(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 11 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type galaxy struct {
  id  int
  row int
  col int
}

// 82000210 too low
func shortestPathGalaxies(lines []string) (int, int) {

  galaxies1 := expandUniverse(lines, 2)
  galaxies2 := expandUniverse(lines, 1000000)

  sum1 := 0
  sum2 := 0
  
  for i := 0; i < len(galaxies1) - 1; i++ {
    for j := i + 1; j < len(galaxies1); j++ {
      rowDiff := abs(galaxies1[i].row - galaxies1[j].row)
      colDiff := abs(galaxies1[i].col - galaxies1[j].col)
      sum1 += rowDiff + colDiff
    }
  }

  for i := 0; i < len(galaxies2) - 1; i++ {
    for j := i + 1; j < len(galaxies2); j++ {
      rowDiff := abs(galaxies2[i].row - galaxies2[j].row)
      colDiff := abs(galaxies2[i].col - galaxies2[j].col)
      sum2 += rowDiff + colDiff
    }
  }
  
  return sum1, sum2
}

func expandUniverse(lines []string, times int) ([]galaxy) {
  var rows []int
  var cols []int
  var galaxyCols []int

  rowsMax := len(lines)
  colsMax := len(lines[0])

  for i := 0; i < rowsMax; i++ {
    isThereGalaxy := false

    for j := 0; j < colsMax; j++ {
      if lines[i][j] == '#' {
        isThereGalaxy = true
        galaxyCols = append(galaxyCols, j)
      }
    }

    if !isThereGalaxy {
      rows = append(rows, i)
    }

    if i == rowsMax - 1 {
      for k := 0; k < colsMax; k++ {
        if !slices.Contains(galaxyCols, k) {
          cols = append(cols, k)
        }
      }
    }
  }

  var galaxies []galaxy
  countGalaxy := 1
  countRow := 0
  countCol := 0

  for i := 0; i < rowsMax; i++ {
    if slices.Contains(rows, i) {
      countRow += times - 1
    }
    
    countCol = 0
    for j := 0; j < colsMax; j++ {
      if slices.Contains(cols, j) {
        countCol += times - 1
      } 
      if string(lines[i][j]) == "#" {
        galaxies = append(galaxies, galaxy{row: i + countRow, col: j + countCol, id: countGalaxy})
        countGalaxy++
      }
    }
  }

  return galaxies
}

func abs(number int) int {
  if number < 0 {
    return number * (-1)
  }
  return number
}
