package solutions

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"time"
)

func Day14() {
  start := time.Now()

  file, err := os.Open("input/day14")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer file.Close()

  reader := bufio.NewReader(file)

  var lines [][]rune
  var line []rune

  for {
    char, _, err := reader.ReadRune()
    if err != nil {
      break
    }

    if char != 10 {
      line = append(line, char)
    } else {
      lines = append(lines, line)
      line = []rune{}
    }
  }

  rollNorth(lines)
  part1 := calcLoad(lines)

  rollWest(lines)
  rollSouth(lines)
  rollEast(lines)

  var seen []string
  var loads []int
  var cycle int

  prev := linesToString(lines)
  loads = append(loads, calcLoad(lines))

  seen = append(seen, prev)
  
  for i := 1; i < 1000000000; i++ {
    rollNorth(lines)
    rollWest(lines)
    rollSouth(lines)
    rollEast(lines)

    prev = linesToString(lines)
    loads = append(loads, calcLoad(lines))

    if slices.Contains(seen, prev) {
      cycle = i
      break
    }
    seen = append(seen, prev)
  }

  cycleStart := slices.Index(seen, prev)
  cycleLen := cycle - cycleStart
  idx := (1000000000 - cycleStart - 1) % cycleLen

  part2 := loads[idx + cycleStart]

	elapsed := time.Since(start)
	fmt.Printf("Day 14 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

func linesToString(lines [][]rune) string {
  res := ""
  for _, line := range lines {
    for _, c := range line {
      res += string(c)
    }
  }
  return res
}

func calcLoad(lines [][]rune) int {
  max := len(lines)
  sum := 0
  for i, line := range lines {
    for _, c := range line {
      if c == 'O' {
        sum += max - i
      }
    }
  }
  return sum
}

func rollNorth(lines [][]rune) {
  var aux rune

  for c := range lines[0] {
    count := 0
    for r := range lines {
      if lines[r][c] == '.' {
        count++
      } else if lines[r][c] == '#' {
        count = 0
      } else if lines[r][c] == 'O' {
        aux = lines[r][c]
        lines[r][c] = lines[r-count][c]
        lines[r-count][c] = aux
        aux = 0
      }
    }
  }
}

func rollWest(lines [][]rune) {
  var aux rune

  for r := range lines {
    count := 0
    for c := range lines[0] {
      if lines[r][c] == '.' {
        count++
      } else if lines[r][c] == '#' {
        count = 0
      } else if lines[r][c] == 'O' {
        aux = lines[r][c]
        lines[r][c] = lines[r][c-count]
        lines[r][c-count] = aux
        aux = 0
      }
    }
  }
}

func rollSouth(lines [][]rune) {
  var aux rune

  for c := range lines[0] {
    count := 0
    for r := len(lines) - 1; r >= 0; r-- {
      if lines[r][c] == '.' {
        count++
      } else if lines[r][c] == '#' {
        count = 0
      } else if lines[r][c] == 'O' {
        aux = lines[r][c]
        lines[r][c] = lines[r+count][c]
        lines[r+count][c] = aux
        aux = 0
      }
    }
  }
}


func rollEast(lines [][]rune) {
  var aux rune

  for r := range lines {
    count := 0
    for c := len(lines[0]) - 1; c >= 0; c-- {
      if lines[r][c] == '.' {
        count++
      } else if lines[r][c] == '#' {
        count = 0
      } else if lines[r][c] == 'O' {
        aux = lines[r][c]
        lines[r][c] = lines[r][c+count]
        lines[r][c+count] = aux
        aux = 0
      }
    }
  }
}


