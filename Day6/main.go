package main

import (
  "io"
  "os"
  "fmt"
  "bufio"
)

var guard = [4]byte{'<','^','>','v'}
var stopped = [4]byte{'l','u','r','d'}
var crate = byte('#')

var dirs = [8][2]int{
  {-1,  0}, // Up
  { 1,  0}, // Down
  { 0, -1}, // Left
  { 0,  1}, // Right
}

func main() {
  var count int
  //mape := getMap()
  mape2 := getMap2()
  rows, cols := len(mape2), len(mape2[0])
  
  //for i := range mape {
  //  fmt.Printf("%s\n", mape[i])
  //}
  for i := range mape2 {
    fmt.Printf("%s", mape2[i])
  }
  max := 16900
  start := 0
  for start != max {
    start++
    if canMove(mape2) && isNotEnd(mape2) {
      move(mape2)
    } else {
      _, pos, face := getGuardFaceDir(mape2)
      for a := 0; a < len(guard)-1 ; a++ {
        if face == guard[a] && face != guard[3]{
          mape2[pos[0]][pos[1]] = guard[a+1]
        } else if face == guard[3] {
          mape2[pos[0]][pos[1]] = guard[0]
        }
      }
      if start <= max {
        continue
      }
    }
  }

  fmt.Printf("\n")
  for i := range mape2 {
    fmt.Printf("%s", mape2[i])
  }

  for row := range rows {
    for col := range cols {
      if mape2[row][col] == byte('x') {
        count++
      }
    }
  }

  fmt.Printf("\nSolution 1: %d\n", count+1)
}

func move(mape [][]byte) {
  nextPos, pos, face  := getGuardFaceDir(mape)
  nextRow := nextPos[0]
  nextCol := nextPos[1]
  row := pos[0]
  col := pos[1]
 
  for i := range guard {
    if mape[row][col] == guard[i] {
      mape[nextRow][nextCol] = face
      mape[row][col] = byte('x')
    }
    if col == 0 {
      mape[row][col] = byte('^')
      break
    }
  }
}

func isNotEnd(mape[][]byte) bool {
  var row, col int

  for i := range mape {
    for j := range mape[i] {
      for a := range guard {
        if mape[i][j] == guard[a] {
          row, col = i, j
        }
      }
    }
  }

  if col == 0 {
    return false
  }

  if row == 0 {
    return false
  }

  return true
}

func canMove(mape [][]byte) bool {
  nextPos, _, _ := getGuardFaceDir(mape)
  nextRow := nextPos[0]
  nextCol := nextPos[1]

  if mape[nextRow][nextCol] == '.' || mape[nextRow][nextCol] == 'x' {
    return true
  }

  if mape[nextRow][nextCol] != crate {
    return true
  }

  return false
}

func getGuardFaceDir(mape [][]byte) ([2]int, [2]int, byte) {
  rows, cols := len(mape), len(mape[0])
  var x, y int
  var nextRow, nextCol int
  var guardFace byte

  for row := range rows {
    for col := range cols {
      if row < 0 || row >= rows ||  col < 0 || col >= cols {
        continue
      }
      //fmt.Printf("analyze: (row: %d) (col: %d)\n", row, col)
      switch mape[row][col] {
      case guard[0]:
        if col-1 >= 0 {
          x = row
          y = col
          nextRow = row
          nextCol = col-1
          guardFace = mape[row][col]
        }
      case guard[1]:
        if row-1 >= 0 {
          x = row
          y = col
          nextRow = row-1
          nextCol = col
          guardFace = mape[row][col]
        }
      case guard[2]:
        if col+1 < cols {
          x = row
          y = col
          nextRow = row
          nextCol = col+1
          guardFace = mape[row][col]
        }
      case guard[3]:
        if row+1 < rows {
          x = row
          y = col
          nextRow = row+1
          nextCol = col
          guardFace = mape[row][col]
        }
      }
    }
  }
  
  return [2]int{nextRow,nextCol}, [2]int{x, y}, guardFace
}

func getMap2() [][]byte {
  var mape [][]byte

  file, err := os.Open("./input.txt")
  if err != nil {
    fmt.Printf("%v", err)
    os.Exit(1)
  }
  defer file.Close()

  reader := bufio.NewReader(file)
  for {
    line, err := reader.ReadBytes('\n')
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Printf("%v", err)
      break
    }
    mape = append(mape, line)
  }
  
  return mape
}
func getMap() [][]byte {
  var mape [][]byte

  file, err := os.Open("./example-input.txt")
  if err != nil {
    fmt.Printf("%v", err)
    os.Exit(1)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    mape = append(mape, scanner.Bytes())
  }

  return mape
}
