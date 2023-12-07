package solutions

import (
	"fmt"
	"strconv"
	"time"
	"unicode"

	"aoc2023/utils"
)

func Day3() {

  start := time.Now()

  lines, err := utils.ReadLines("input/day3.txt")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  part1 := engineNumbers(lines)
  part2 := gearRatios(lines)

  elapsed := time.Since(start)
  fmt.Printf("Day 03 part1 %v\t", part1)
  fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type number struct {
  line  int
  start int
  end   int
  val   int
}

type symbol struct {
  line  int
  col   int
}

func engineNumbers(input []string) int {
  sum := 0
  symbols, numbers := getSymbolsAndNumbers(input, false)

  for _, num := range numbers {
    for _, sym := range symbols {
      lineClose := num.line - sym.line
      if lineClose < 2 && lineClose > -2 {
        colStart := num.start - sym.col
        colEnd := num.end - sym.col
        if colStart < 2 && colStart > -2 {
          sum += num.val
        } else if colEnd < 2 && colEnd > -2 {
          sum += num.val
        }
      }
    }
  }

  return sum
}

func gearRatios(input []string) int {
  sum := 0
  symbols, numbers := getSymbolsAndNumbers(input, true)

  for _, sym := range symbols {
    num1 := 0
    num2 := 0
    for _, num := range numbers {
      line := num.line - sym.line

      if line < 2 && line > -2 {
        colS := num.start - sym.col
        colE := num.end - sym.col
        
        if colS < 2 && colE > -2 {
          if num1 == 0 {
            num1 = num.val
          } else {
            num2 = num.val
          }
        }
      }
    }
    sum += num1 * num2
  }

  return sum
}

func isSymbol(r rune, part2 bool) bool {
  if part2 {
    return r == '*'
  }
  return !unicode.IsDigit(r) && r != '.'
}

func getSymbolsAndNumbers(input []string, part2 bool) ([]symbol, []number) {

  var symbols []symbol
  var numbers []number

  numStart := -1
  numEnd := -1
  prevNum := false
  val := ""

  for i, line := range input {
    for j, char := range line {
      if isSymbol(char, part2) {
        symbols = append(symbols, symbol{line: i, col: j})
      }

      if unicode.IsDigit(char) {
        if !prevNum {
          numStart = j
          prevNum = true
        }
        if prevNum {
          val += string(char)
          numEnd = j
        }
      } else {
        if prevNum {
          prevNum = false
          numVal, _ := strconv.Atoi(val)
          val = ""
          numbers = append(numbers, 
          number{ line: i, start: numStart, end: numEnd, val: numVal, })
        }
      }
    }
  }

  return symbols, numbers
}
