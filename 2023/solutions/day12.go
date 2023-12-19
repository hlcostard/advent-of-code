package solutions

import (
	"fmt"
	"slices"
	"strconv"
	"time"

	"aoc2023/utils"
)

func Day12() {
	start := time.Now()

	lines, err := utils.ReadRunes("input/day12")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1, part2 := hotSpringArrangement(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 12 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

var cache = make(map[string]int)

func hotSpringArrangement(lines [][]rune) (part1 int, part2 int) {

  sum1 := 0
  sum2 := 0

  for _, line := range lines {

    split := 100
    var num string
    var number int
    var combinations []int

    for i := 0; i < len(line); i++ {
      if line[i] == ' ' {
        split = i
      }
      if split < i {
        if line[i] >= '0' && line[i] <= '9' {
          num += string(line[i])
        } else if line[i] == ',' {
          number, _ = strconv.Atoi(num)
          combinations = append(combinations, number)
          num = "" 
        }
      }
    }
    number, _ = strconv.Atoi(num)
    combinations = append(combinations, number)

    springs := line[:split]
    springs2 := springs
    combinations2 := combinations

    for i := 1; i < 5; i++ {
      springs2 = append(springs2, '?')
      springs2 = append(springs2, springs...)
      combinations2 = append(combinations2, combinations...)
    }

    sum1 += springCheck(springs, combinations)
    sum2 += springCheck(springs2, combinations2)
  }

  return sum1, sum2
}

func springCheck(springs []rune, combinations []int) int {

  key := string(springs)

  for _, number := range combinations {
    key += strconv.Itoa(number) + ","
  }

  if v, ok := cache[key]; ok {
    return v
  }
  

  // remove leading dots
  minIdx := minIndex(slices.Index(springs, '#'),slices.Index(springs, '?'))
  s := springs[minIdx:]

  if len(s) == 0 {
    if len(combinations) == 0 {
      cache[key] = 1
      return 1
    } else {
      cache[key] = 0
      return 0
    }
  }

  if len(combinations) == 0 {
    if !slices.Contains(s, '#') {
      cache[key] = 1
      return 1
    } else {
      cache[key] = 0
      return 0
    }
  }

  if s[0] == '#' {
    if len(s) < combinations[0] || slices.Contains(s[:combinations[0]], '.') {
      cache[key] = 0
      return 0
    }

    if len(s) == combinations[0] {
      if len(combinations) == 1 {
        cache[key] = 1
        return 1
      } else {
        cache[key] = 0
        return 0
      }
    }

    if s[combinations[0]] == '#' {
      cache[key] = 0
      return 0
    }
    res := springCheck(s[combinations[0]+1:], combinations[1:])
    cache[key] = res
    return res
  }

  if s[0] == '?' {
    newS := append([]rune{'#'}, s[1:]...)
    res := springCheck(newS, combinations) + springCheck(s[1:], combinations)
    cache[key] = res
    return res
  }
  
  res := springCheck(s[1:], combinations)
  cache[key] = res
  return res
}

func minIndex(num1, num2 int) int {
  if num1 == -1 {
    if num2 == -1 {
      return 0
    }
    return num2
  } 
  if num2 == -1 {
    return num1
  }
  if num1 < num2 {
    return num1
  }
  return num2
}
