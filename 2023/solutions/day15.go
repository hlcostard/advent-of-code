package solutions

import (
	"fmt"
	"time"
	"unicode"

	"aoc2023/utils"
)

func Day15() {
	start := time.Now()

	lines, err := utils.ReadRunes("input/day15")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1 := initSequence(lines)
  part2 := focusingPower(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 15 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

func initSequence(lines [][]rune) int {
  
  res := 0
  var input []rune

  for _, c := range lines[0] {
    if c != ',' {
      input = append(input, c)
    } else {
      res += hash(input)
      input = []rune{}
    }
  }
  res += hash(input)

  return res 
}

type box struct {
  lenses  []lens  
}

type lens struct {
  label   string
  length  int
}

func focusingPower(lines [][]rune) int {
  hashmap := lines[0]
  var label []rune
  var boxes [256]box

  i := 0
  for i < len(hashmap) {
    if unicode.IsDigit(hashmap[i]) {
      i++
      continue
    }
    if hashmap[i] != ',' && hashmap[i] != '=' && hashmap[i] != '-' {
      label = append(label, hashmap[i])

    } else {
      pos := hash(label)

      if hashmap[i] == '=' {
        length := int(hashmap[i+1] - '0')
        addLens(string(label), pos, length, &boxes)
        i++

      } else if hashmap[i] == '-' {
        removeLens(string(label), pos, &boxes)
      }
      label = []rune{}
    }
    i++
  }

  res := 0
  for i := range boxes {
    box := i + 1
    if len(boxes[i].lenses) > 0 {
      for j, l := range boxes[i].lenses {
        slot := j + 1
        res += box * slot * l.length
      }
    }
  }

  return res
}

func addLens(label string, pos, length int, boxes *[256]box) {
  index := -1
  for i, l := range boxes[pos].lenses {
    if l.label == label {
      index = i
      break
    }
  }

  if index == -1 {
    boxes[pos].lenses = append(boxes[pos].lenses, lens{label: label, length: length})
    return
  }

  l := lens{label: label, length: length} 
  boxes[pos].lenses[index] = l
}

func removeLens(label string, pos int, boxes *[256]box) {
  index := -1
  for i, l := range boxes[pos].lenses {
    if l.label == label {
      index = i
      break
    }
  }

  if index > -1 {
    boxes[pos].lenses = append(boxes[pos].lenses[:index], boxes[pos].lenses[index+1:]...)
  }
}

func hash(input []rune) int {
  res := 0

  for _, c := range input {
    res += int(c)
    res *= 17
    res = res % 256
  }

  return res
}
