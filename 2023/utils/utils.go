package utils

import (
	"bufio"
	"os"
)

func ReadLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    
    return lines, scanner.Err()
}

func ReadRunes(path string) ([][]rune, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines [][]rune
  var line []rune

  reader := bufio.NewReader(file)
  for {
    char, _, err := reader.ReadRune()
    if err != nil {
      break
    }

    if char != 10 {
      line = append(line, char)
    } else {
      lines = append(lines, line)
      line = []rune{}
    }
  }

  return lines, nil
}

func RunesToInt (input [][]rune) [][]int {
  var lines [][]int
  var line []int
  for i := range input {
    for j := range input {
      line = append(line, int(input[i][j] - '0'))
    }
    lines = append(lines, line)
    line = []int{}
  }

  return lines
}
