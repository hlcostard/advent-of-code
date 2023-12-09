package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"aoc2023/utils"
)

func Day9() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day9.txt")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1, part2 := predictValue(lines)

	elapsed := time.Since(start)
	fmt.Printf("Day 09 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

func predictValue(input []string) (int, int) {

  next := 0
  prev := 0
  for _, line := range input {
    fields := strings.Fields(line)

    var numbers []int
    for _, field  := range fields {
      num, _ := strconv.Atoi(field)
      numbers = append(numbers, num)
    }

    n, p := stepsDiff(numbers)

    next += numbers[len(numbers)-1] + n
    prev += numbers[0] - p
  }

  return next, prev
}

func stepsDiff(nums []int) (int, int) {
  var nums2 []int
  for i := 0; i < len(nums) - 1; i++ {
    diff := nums[i+1] - nums[i]
    nums2 = append(nums2, diff)
  }

  if nums2[len(nums2) - 1] != 0 {
    add, diff := stepsDiff(nums2)
    return nums2[len(nums2) - 1] + add, nums2[0] - diff
  }

  return nums2[len(nums2) - 1], nums2[0]
}
