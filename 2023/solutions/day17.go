package solutions

import (
	"fmt"
	"time"

	"aoc2023/utils"
)

func Day17() {
	start := time.Now()

	lines, err := utils.ReadRunes("test-input/day17")
	if err != nil {
		fmt.Printf("error reading file %v\n", err)
	}

  nums := utils.RunesToInt(lines)

  part1 := findEnd(nums)
  part2 := 420

	elapsed := time.Since(start)
	fmt.Printf("Day 17 part1 %v\t", part1)
	fmt.Printf("part2 %v\ttime %v \n", part2, elapsed)
}

type block struct {
  row       int
  col       int
  vRow      int
  vCol      int
  dist      int
}

var blocks = make(map[int][]block)
var seen_blocks = make(map[int]block)

func findEnd(layout [][]int) int {
  moveCrucible(0, 0, 0, 1, 0, 1, layout)
  moveCrucible(0, 0, 0, 0, 1, 1, layout)

  for {
    fmt.Println(blocks)
    current := getMinNextHeatLoss()


    fmt.Println(current)
    break

  }
  return 10
}

func moveCrucible(heatLoss, row, col, vRow, vCol, dist int, layout [][]int) {
  maxRow := len(layout) - 1
  maxCol := len(layout[0]) - 1

  row += vRow
  col += vCol

  if row < 0 || col < 0 { 
    return
  }
  if row > maxRow || col > maxCol {
    return
  }

  newHLoss := heatLoss + layout[row][col]

  b := block{row, col, vRow, vCol, dist}

  if _, ok := blocks[]; !ok {
    blocks[b] = newHLoss
    seen_blocks[b] = newHLoss
  }
}

func getMinNextHeatLoss() int {
  min := 10000000
  for i := range blocks {
    if blocks[i] < min {
      min = blocks[i]
    }
  }
  return min
}


// // prevBlock="i,j"
// type block struct {
//   visited   bool
//   row       int
//   col       int
//   shortDist int
//   heatLoss  int
//   prevBlock string
// }
//
// func moveCrucible(layout [][]int) int {
//
//
//   blocks := make([][]block, len(layout))
//   for i := range blocks {
//     blocks[i] = make([]block, len(layout))
//     for j := range blocks[i] {
//       blocks[i][j].row = i
//       blocks[i][j].col = j
//       blocks[i][j].visited = false
//       blocks[i][j].heatLoss = layout[i][j]
//       blocks[i][j].shortDist = 1000000000
//     }
//   }
//   blocks[0][0].visited = true
//   blocks[0][0].shortDist = 0
//
//   i := 0
//   j := 0
//
//   for !visitedAll(blocks) {
//     i, j = checkNext(i, j, blocks)
//     blocks[i][j].visited = true
//   }
//
//   maxRow := len(blocks) - 1
//   maxCol := len(blocks[0]) - 1
//
//   return blocks[maxRow][maxCol].shortDist
// }
//
// func countBlocks(blocks [][]block) (int, int) {
//   t := 0
//   f := 0
//   for i := range blocks {
//     for j := range blocks[i] {
//       if blocks[i][j].visited {
//         t++
//       } else {
//         f++
//       }
//     }
//   }
//   return t, f
// }
//
// func visitedAll(blocks [][]block) bool {
//   for i := range blocks {
//     for j := range blocks[i] {
//       if !blocks[i][j].visited {
//         return false
//       }
//     }
//   }
//   return true
// }
//
//
// func checkNext(i, j int, blocks [][]block) (int, int) {
//   if i > 0 { 
//     if !blocks[i-1][j].visited {
//       if blocks[i-1][j].shortDist > blocks[i-1][j].heatLoss + blocks[i][j].shortDist {
//         blocks[i-1][j].shortDist = blocks[i-1][j].heatLoss + blocks[i][j].shortDist
//         blocks[i-1][j].prevBlock = fmt.Sprintf("%d,%d", i, j)
//       }
//     }
//   }
//
//   if i < len(blocks) - 1 {
//     if !blocks[i+1][j].visited {
//       if blocks[i+1][j].shortDist > blocks[i+1][j].heatLoss + blocks[i][j].shortDist {
//         blocks[i+1][j].shortDist = blocks[i+1][j].heatLoss + blocks[i][j].shortDist
//         blocks[i+1][j].prevBlock = fmt.Sprintf("%d,%d", i, j)
//       }
//     }
//   }
//
//   if j > 0 {
//     if !blocks[i][j-1].visited {
//       if blocks[i][j-1].shortDist > blocks[i][j-1].heatLoss + blocks[i][j].shortDist {
//         blocks[i][j-1].shortDist = blocks[i][j-1].heatLoss + blocks[i][j].shortDist
//         blocks[i][j-1].prevBlock = fmt.Sprintf("%d,%d", i, j)
//       }
//     }
//   }
//
//   if j < len(blocks[0]) - 1 {
//     if !blocks[i][j+1].visited {
//       if blocks[i][j+1].shortDist > blocks[i][j+1].heatLoss + blocks[i][j].shortDist {
//         blocks[i][j+1].shortDist = blocks[i][j+1].heatLoss + blocks[i][j].shortDist
//         blocks[i][j+1].prevBlock = fmt.Sprintf("%d,%d", i, j)
//       }
//     }
//   }
//
//   var minI int
//   var minJ int
//   min := 1000000000
//   for i := range blocks {
//     for j := range blocks[0] {
//       if blocks[i][j].shortDist < min && !blocks[i][j].visited {
//         min = blocks[i][j].shortDist
//         minI = i
//         minJ = j
//       }
//     }
//   }
//
//   return minI, minJ
// }
