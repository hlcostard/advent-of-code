package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc2023/utils"
)

func Day6() {
  start := time.Now()

  lines, err := utils.ReadLines("input/day06")
  if err != nil {
    fmt.Printf("error reading file %v\n", err)
  }
  _ = lines

  part1 := beatRecords(lines)
  part2 := actuallyItsJustBeatRecord(lines)

  elapsed := time.Since(start)
  fmt.Printf("Day 06 part1 %v\t", part1)
  fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

// part1
func beatRecords(input []string) int {

  var times []int
  var distances []int

  for i := 1; i < len(strings.Fields(input[0])); i++ {
    time, _ := strconv.Atoi(strings.Fields(input[0])[i])
    distance, _ := strconv.Atoi(strings.Fields(input[1])[i])
    times = append(times, time)
    distances = append(distances, distance)
  }

  result := 1
  for i := range times {
    result *= countTimesCanBeatRecord(times[i], distances[i])
  }

  return result
}

func actuallyItsJustBeatRecord(input []string) int {

  time := ""
  distance := ""

  for i := 1; i < len(strings.Fields(input[0])); i++ {
    time = time + strings.Fields(input[0])[i]
    distance = distance + strings.Fields(input[1])[i]
  }

  timeInt, _ := strconv.Atoi(time)
  distanceInt, _ := strconv.Atoi(distance)
  
  return countTimesCanBeatRecord(timeInt, distanceInt)
}

func countTimesCanBeatRecord(totalTime, distanceRecord int) int {
  count := 0
  distance := 0
  for speed := 1; speed < totalTime; speed++ {
    time := totalTime - speed 
    distance = speed * (time)
    if distance > distanceRecord {
      count++
    }
  }

  return count
}


