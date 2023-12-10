package solutions

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "time"
)

func Day1() {

  start := time.Now()
  sum := 0
  sum2 := 0

  file, err := os.Open("input/day1.txt")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    sum += getNumbers(line, 1)
    sum2 += getNumbers(line, 2)
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  elapsed := time.Since(start)
  fmt.Printf("Day 01 part1 %v\t", sum)
  fmt.Printf("part2 %v \ttime %v \n", sum2, elapsed)
}

func getNumbers(input string, part int) int {

  // comment the line with the numbers name to solve part1
  numbersPart1 := [10]string{ 
    "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
  }
  numbersPart2 := [20]string{ 
    "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
  }

  var numbers []string

  for i := 0; i < 10 * part; i++ {
    if part == 1 {
      numbers = append(numbers, numbersPart1[i])
    } else {
      numbers = append(numbers, numbersPart2[i])
    }
  }


  var first int
  var last int
  firstId := 100000
  lastId := -10

  for i, number := range numbers {
    index := strings.Index(input, number)
    lastIndex := strings.LastIndex(input, number)

    if index > -1 {
      if index < firstId {
        firstId = index
        if i < 10 {
          first = i
        } else {
          first = i - 10
        }
      }

      if lastIndex > lastId {
        lastId = lastIndex
        if i < 10 {
          last = i
        } else {
          last = i - 10
        }
      }
    }
  }

  return first * 10 + last 
}
