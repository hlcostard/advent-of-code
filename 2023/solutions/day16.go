package solutions

import (
	"fmt"
	"time"

	"aoc2023/utils"
)

func Day16() {
	start := time.Now()

	lines, err := utils.ReadRunes("input/day16")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

  part1 := countEnergized(lines, 0, 0, "right")
  part2 := maxEnergy(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 16 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type beam struct {
  r bool
  l bool
  u bool
  d bool
}

func countEnergized(layout [][]rune, i, j int, direction string) int {

  beams := make([][]beam, len(layout))
  energize := make([][]rune, len(layout))
  for i := range energize {
    energize[i] = make([]rune, len(layout[0]))
    beams[i] = make([]beam, len(layout[0]))
    for j := range energize[i] {
      energize[i][j] = '.'
      beams[i][j].r = false
      beams[i][j].l = false
      beams[i][j].u = false
      beams[i][j].d = false
    }
  }

  switch direction {
  case "right":
    moveRight(energize, layout, i, j, beams)
  case "left":
    moveLeft(energize, layout, i, j, beams)
  case "up":
    moveUp(energize, layout, i, j, beams)
  case "down":
    moveDown(energize, layout, i, j, beams)
  }

  r := 0
  for _, line := range energize {
    for _, c := range line {
      if c == '#' {
        r++
      }
    }
  }

  return r
}

func moveRight(energize, layout [][]rune, i, j int, beams [][]beam) {
  if j >= len(energize[0]) {
    return
  }
  if beams[i][j].r {
    return
  }

  energize[i][j] = '#'
  beams[i][j].r = true

  switch layout[i][j] {
  case '.':
    moveRight(energize, layout, i, j+1, beams)

  case '|':
    moveUp(energize, layout, i-1, j, beams)
    moveDown(energize, layout, i+1, j, beams)

  case '-':
    moveRight(energize, layout, i, j+1, beams)

  case '/':
    moveUp(energize, layout, i-1, j, beams)

  case '\\':
    moveDown(energize, layout, i+1, j, beams)
  }
}

func moveLeft(energize, layout [][]rune, i, j int, beams [][]beam) {
  if j < 0 {
    return
  }
  if beams[i][j].l {
    return
  }

  energize[i][j] = '#'
  beams[i][j].l = true

  switch layout[i][j] {
  case '.':
    moveLeft(energize, layout, i, j-1, beams)

  case '|':
    moveUp(energize, layout, i-1, j, beams)
    moveDown(energize, layout, i+1, j, beams)

  case '-':
    moveLeft(energize, layout, i, j-1, beams)

  case '/':
    moveDown(energize, layout, i+1, j, beams)

  case '\\':
    moveUp(energize, layout, i-1, j, beams)
  }
}

func moveUp(energize, layout [][]rune, i, j int, beams [][]beam) {
  if i < 0 {
    return
  }
  if beams[i][j].u {
    return
  }

  energize[i][j] = '#'
  beams[i][j].u = true

  switch layout[i][j] {
  case '.':
    moveUp(energize, layout, i-1, j, beams)

  case '|':
    moveUp(energize, layout, i-1, j, beams)

  case '-':
    moveRight(energize, layout, i, j+1, beams)
    moveLeft(energize, layout, i, j-1, beams)

  case '/':
    moveRight(energize, layout, i, j+1, beams)

  case '\\':
    moveLeft(energize, layout, i, j-1, beams)
  }
}

func moveDown(energize, layout [][]rune, i, j int, beams [][]beam) {
  if i >= len(energize) {
    return
  }
  if beams[i][j].d {
    return
  }

  energize[i][j] = '#'
  beams[i][j].d = true

  switch layout[i][j] {
  case '.':
    moveDown(energize, layout, i+1, j, beams)

  case '|':
    moveDown(energize, layout, i+1, j, beams)

  case '-':
    moveRight(energize, layout, i, j+1, beams)
    moveLeft(energize, layout, i, j-1, beams)

  case '/':
    moveLeft(energize, layout, i, j-1, beams)

  case '\\':
    moveRight(energize, layout, i, j+1, beams)
  }
}

func maxEnergy(layout [][]rune) int {
  var res []int
  maxI := len(layout) - 1
  maxJ := len(layout[0]) - 1

  for i := 0; i < len(layout); i++ {
    res = append(res, countEnergized(layout, i, 0, "right"))
    res = append(res, countEnergized(layout, i, maxJ, "left"))
    res = append(res, countEnergized(layout, maxI, i, "up"))
    res = append(res, countEnergized(layout, 0, i, "down"))
  }

  return maxInt(res)
}

func maxInt(input []int) int {
  res := 0
  for i := range input {
    if input[i] > res {
      res = input[i]
    }
  }

  return res
}
