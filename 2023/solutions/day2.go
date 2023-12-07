package solutions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func Day2() {

  start := time.Now()

  file, err := os.Open("input/day2.txt")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }
  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  sum := 0
  powerSum := 0

  for scanner.Scan() {
    line := scanner.Text()

    sum = sum + possibleGame(line)
    powerSum = powerSum + powerCubes(line)
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  elapsed := time.Since(start)
  fmt.Printf("Day 02 part1 %v\t", sum)
  fmt.Printf("part2 %v \ttime %v \n", powerSum, elapsed)
}

// part 1
func possibleGame(line string) int {
  max := map[string]int{
    "red": 12,
    "green": 13,
    "blue": 14,
  }

  fields := strings.FieldsFunc(line, splitDay2)

  id, _ := strconv.Atoi(fields[1])

  i := 3
  for i < len(fields) {
    num, _ := strconv.Atoi(fields[i-1])

    if num > max[fields[i]] {
      return 0
    }

    i = i + 2
  }

  return id
}

// part 2
func powerCubes (line string) int {
  max := map[string]int{
    "red": 0,
    "green": 0,
    "blue": 0,
  }

  fields := strings.FieldsFunc(line, splitDay2)

  i := 3
  for i < len(fields) {
    num, _ := strconv.Atoi(fields[i-1])

    if num > max[fields[i]] {
      max[fields[i]] = num
    }

    i = i + 2
  }

  return max["red"] * max["blue"] * max["green"]
}


// split the input 
func splitDay2(c rune) bool {
  return c == ':' || c == ',' || c == ';' || c == ' '
}


