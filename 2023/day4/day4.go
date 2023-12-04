package day4

import (
	"aoc2023/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"
)

func Solve() {

  start := time.Now()

  lines, err := utils.ReadLines("input/day4")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  part1 := winningNumbers(lines)
  part2 := winningCards(lines)

  elapsed := time.Since(start)
  fmt.Printf("Day 04 part 1 %v\t", part1)
  fmt.Printf("part 2 %v\ttime: %v \n", part2, elapsed)
}

func winningNumbers(input []string) int {

  points := 0

  for _, line := range input {
    fields := strings.FieldsFunc(line, split)
    winning := strings.Fields(fields[1])
    numbers := strings.Fields(fields[2])

    cardPoints := 0

    for _, num := range numbers {
      if slices.Contains(winning, num) {
        if cardPoints == 0 {
          cardPoints = 1
        } else {
          cardPoints = cardPoints * 2
        }
      }
    }

    points += cardPoints
  }

  return points
}

type card struct {
  cardNum string
  amount  int
}

func winningCards(input []string) int {

  cards := []card{}
  totalCards := 0

  for i := 1; i <= len(input); i++ {
    cards = append(cards, card{ cardNum: strconv.Itoa(i), amount: 1})
  }

  for i, line := range input {
    fields := strings.FieldsFunc(line, split)
    winning := strings.Fields(fields[1])
    numbers := strings.Fields(fields[2])

    count := 0

    for _, num := range numbers {
      if slices.Contains(winning, num) {
        count++
        cards[i+count].amount = cards[i+count].amount + 1 * cards[i].amount
      }
    }
    totalCards += cards[i].amount
  }

  return totalCards
}

func split(c rune) bool {
  return c == ':' || c == '|' 
}

