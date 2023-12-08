package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc2023/utils"
)

func Day7() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day7.txt")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}
	_ = lines

	part1 := camelCards(lines, 1)
	part2 := camelCards(lines, 2)

	elapsed := time.Since(start)
	fmt.Printf("Day 07 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type hand struct {
	hand     string
	handType string
	bid      int
	rank     int
}

func camelCards(input []string, part int) int {
	var hands []hand

	for _, line := range input {
		fields := strings.Fields(line)
		bid, _ := strconv.Atoi(fields[1])

    handT := rankHand(fields[0], part) 
		hands = append(hands, hand{hand: fields[0], bid: bid, handType: handT})
	}

  result := 0
  for i := range hands {
    count := 1
    for j := range hands {
      if i != j && compareHands(hands[i], hands[j], part) {
        count++
      }
    }
    hands[i].rank = count
    result += hands[i].bid * count
  }

	return result
}

func compareHands(hand1, hand2 hand, part int) bool {

	cardRank := map[string]int{"2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7,
  "8": 8, "9": 9, "T": 10, "J": 11, "Q": 12, "K": 13, "A": 14}

  cardRank2 := map[string]int{"J": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, 
  "7": 7, "8": 8, "9": 9, "T": 10, "Q": 12, "K": 13, "A": 14}

	handType := map[string]int{"HC": 1, "2K": 2, "TP": 3, "3K": 4, "FH": 5,
  "4K": 6, "5K": 7}

  if handType[hand1.handType] == handType[hand2.handType] {
    for i := 0; i < 5; i++ {
      if part == 1 {
        if cardRank[string(hand1.hand[i])] > cardRank[string(hand2.hand[i])] {
          return true
        } 
        if cardRank[string(hand1.hand[i])] < cardRank[string(hand2.hand[i])] {
          return false
        }
      } else {
        if cardRank2[string(hand1.hand[i])] > cardRank2[string(hand2.hand[i])] {
          return true
        } 
        if cardRank2[string(hand1.hand[i])] < cardRank2[string(hand2.hand[i])] {
          return false
        }
      }
    }
  }
  return handType[hand1.handType] > handType[hand2.handType] 
}

func rankHand(hand string, part int) string {
	count := []int{1, 1, 1, 1, 1}
  countJ := 0

  if part == 1 {
    for i := 0; i < 5; i++ {
      for j := 0; j < 5; j++ {
        if i != j && string(hand[i]) == string(hand[j]) {
          count[i]++
        }
      }
    }
  } else {
    for i := 0; i < 5; i++ {
      if string(hand[i]) == "J" {
        countJ++
      } else {
        for j := 0; j < 5; j++ {
          if i != j && string(hand[i]) == string(hand[j]) {
            count[i]++
          }
        }
      }
    }
  }

	max := max(count)
  handT := "HC"

	if max == 5 || countJ == 5 {
		handT = "5K"
	} else if max == 4 {
    if countJ == 1 {
      handT = "5K"
    } else {
      handT = "4K"
    }
	} else if max == 3 {
    if contains(count, 2) {
      handT = "FH"
    } else if countJ == 2 {
      handT ="5K"
    } else if countJ == 1 {
      handT = "4K"
    } else {
      handT = "3K"
    }
  } else if max == 2 {
    if countNum(count, 2) == 4 {
      if countJ == 1 {
        handT = "FH"
      } else {
        handT = "TP"
      }
    } else {
      if countJ == 3 {
        handT = "5K"
      } else if countJ == 2 {
        handT = "4K"
      } else if countJ == 1 {
        handT = "3K"
      } else {
        handT = "2K"
      }
    }
  } else {
    if countJ == 4 {
      handT = "5K"
    } else if countJ == 3 {
      handT = "4K"
    } else if countJ == 2 {
      handT = "3K"
    } else if countJ == 1 {
      handT = "2K"
    }
  }

	return handT
}

func max(input []int) int {
	max := 0
	for _, num := range input {
		if num > max {
			max = num
		}
	}
	return max
}

func contains(array []int, number int) bool {
  for _, num := range array {
    if num == number {
      return true
    }
  }
  return false
}

func countNum(array []int, number int) int {
  count := 0
  for _, num := range array {
    if num == number {
      count++ 
    }
  }
  return count
}
