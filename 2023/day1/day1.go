package day1

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "time"
)

func Solve() {

  start := time.Now()
  sum := 0

  file, err := os.Open("day1/input2.txt")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  defer file.Close()
  
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    line := scanner.Text()
    sum = sum + getNumbers(line)
  }

  if err := scanner.Err(); err != nil {
    fmt.Printf("error reading file %v\n", err)
  }

  elapsed := time.Since(start)
  fmt.Printf("the result is %v, time: %v \n", sum, elapsed)
}

func getNumbers(input string) int {

  // comment the line with the numbers name to solve part1
  numbers := [20]string{ 
    "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
    "zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
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
