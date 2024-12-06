package main

import (
  "os"
  "fmt"
  "bufio"
)

var directions = [8][2]int {
  {-1,  0}, // Up
  { 1,  0}, // Down
  { 0, -1}, // Left
  { 0,  1}, // Right
  {-1, -1}, // Up-Left Diagonal
  {-1,  1}, // Up-Right Diagonal
  { 1, -1}, // Down-Left Diagonal
  { 1,  1}, // Down-Right Diagonal
}

var diagonals = [4][2]int {
  {-1, -1}, // Up-Left Diagonal
  { 1,  1}, // Down-Right Diagonal
  {-1,  1}, // Up-Right Diagonal
  { 1, -1}, // Down-Left Diagonal
}


func main() {
  fmt.Printf("Solution 1: %d\n", countXMAS(getInput()))
  fmt.Printf("Solution 2: %d\n", countMAS(getInput()))
}

func countMAS(input [][]rune) int {
  count := 0

  for row := 0; row < len(input); row++ {
    for col := 0; col < len(input[row]); col++ {
      if checkMAS(input, row, col) {
        count++
      }
    }
  }
  return count
}

func countXMAS(input [][]rune) int {
  count := 0
  rows, cols := len(input), len(input[0])

  for row := range rows {
    for col := range cols {
      for _, direction := range directions {
        if checkXMAS(input, row, col, direction) {
          count++
        }
      }
    }
  }
  return count
}

func checkMAS(input [][]rune, row, col int) bool {
  rows, cols := len(input), len(input[0])

  if input[row][col] != 'A' {
    return false
  }
  
  var chars []rune
  for _, pos := range diagonals {
    nextRow := row + pos[0]
    nextCol := col + pos[1]
    if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
      return false
    }
    chars = append(chars, input[nextRow][nextCol])
  }

  return (chars[0] == 'M' && chars[2] == 'M' && chars[1] == 'S' && chars[3] == 'S') ||
         (chars[0] == 'S' && chars[2] == 'S' && chars[1] == 'M' && chars[3] == 'M') ||
         (chars[0] == 'M' && chars[3] == 'M' && chars[1] == 'S' && chars[2] == 'S') ||
         (chars[0] == 'S' && chars[3] == 'S' && chars[1] == 'M' && chars[2] == 'M')
}

func checkXMAS(input [][]rune, row, col int, direction [2]int) bool {
  pattern := []rune{'X','M','A','S'}
  rows, cols := len(input), len(input[0])

  for i := range pattern {
    nextRow := row+direction[0]*i
    nextCol := col+direction[1]*i

    if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
      return false
    }

    if input[nextRow][nextCol] != pattern[i] {
      return false
    }
  }
  return true
}

func getInput() [][]rune {
  var list [][]rune
  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Println(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    list = append(list, []rune(scanner.Text()))
  }

  return list
}
