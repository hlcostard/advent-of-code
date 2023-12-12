package solutions

import (
	"fmt"
	"strings"
	"time"

	"aoc2023/utils"
)

func Day8() {
	start := time.Now()

	lines, err := utils.ReadLines("input/day08")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

	part1 := stepsToZ(lines, 1)
	part2 := stepsToZ(lines, 2)

	elapsed := time.Since(start)
	fmt.Printf("Day 08 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type node struct {
  element string
  left    string
  right   string
}

func stepsToZ(input []string, part1 int) int {

  directions := input[0]
  var nodes []node

  for i := 2; i < len(input); i++ {
    fields := strings.Fields(input[i])
    
    nodes = append(nodes, node{
      element: fields[0],
      left: fields[2][1:4],
      right: fields[3][0:3],
    })
  }

  if part1 == 1 {
    currentNode := getNode(nodes, "AAA")
    steps := 0
    i := 0
    for currentNode.element != "ZZZ" {
      if i >= len(directions) {
        i = 0
      }

      next := nextNode(currentNode, string(directions[i]))
      currentNode = getNode(nodes, next)
      steps++
      i++
    }

    return steps
  }

  currentNodes := findInitialNodes(nodes)
  var steps []int
  for id := range currentNodes {
    step := 0
    i := 0
    currentNode := currentNodes[id]

    for string(currentNode.element[2]) != "Z" {
      if i >= len(directions) {
        i = 0
      }

      next := nextNode(currentNode, string(directions[i]))
      currentNode = getNode(nodes, next)
      step++
      i++
    }
    steps = append(steps, step)
  }

  return lcm(steps)
}

func nextNode(node node, direction string) string {
  if direction == "R" {
    return node.right
  } 
  return node.left
}

func getNode(nodes []node, name string) node {
  for _, n := range nodes {
    if n.element == name {
      return n
    }
  }
  return node{}
}

func findInitialNodes(nodes []node) []node {
  var initNodes []node
  for _, n := range nodes {
    if string(n.element[2]) == "A" {
      initNodes = append(initNodes, n)
    }
  }
  return initNodes
}

func gcd(num1, num2 int) int {
  for num2 != 0 {
    aux := num2
    num2 = num1 % num2
    num1 = aux
  }
  return num1
}

func lcm(num []int) int {
  result := num[0] * num[1] / gcd(num[0], num[1])

  for i := 2; i < len(num); i++ {
    result = lcm([]int{result, num[i]})
  }

  return result
}
