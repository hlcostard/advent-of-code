package utils

import (
    "os"
    "bufio"
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

  reader := bufio.NewReader(file)

  var lines [][]rune
  var line []rune

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




