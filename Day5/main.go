package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  queues := queueParser()
  rules := getRules()
  solution1, solution2 := 0, 0

  for _, queue := range queues {
    midNum := getValidMidNum(rules, queue)
    midNum2 := getInvalidMidNum(rules, queue)
    fmt.Printf("Update queue %v: %v\n", queue, midNum)
    fmt.Printf("Update2 queue %v: %v\n", queue, midNum2)
    solution1 += midNum
    solution2 += midNum2
  }

  fmt.Printf("\nSolution 1: %v\n", solution1)
  fmt.Printf("\nSolution 2: %v\n", solution2)
}

func getInvalidMidNum(rules [][2]int, queue []int ) int {
  length := len(queue)
  mid := length / 2

  if !queueIsSorted(rules, queue) {
    sortQueue(rules, queue)
    return queue[mid]
  }

  return 0
}

func getValidMidNum(rules [][2]int, queue []int) int {
  length := len(queue)
  mid := length / 2

  if queueIsSorted(rules, queue) {
    if length % 2 != 0 {
      return queue[mid]
    }
  }
  return 0
}

func sortQueue(rules [][2]int, queue []int) {
  for i := 0; i < len(queue)-1; i++ {
    for j := 0; j < len(queue)-1-i; j++ {
      if !isOrdered(rules, queue[j], queue[j+1]) {
        queue[j], queue[j+1] = queue[j+1], queue[j]
      }
    }
  }
}

func queueIsSorted(rules [][2]int, queue []int) bool {
  for i := 0; i < len(queue)-1; i++ {
    if !isOrdered(rules, queue[i], queue[i+1]) {
      return false
    }
  }
  return true
}

func isOrdered(rules [][2]int, a, b int) bool {
  for _, rule := range rules {
    if int(rule[0]) == a && int(rule[1]) == b {
      return true
    }
  }
  return false
}

func queueParser() [][]int {
  var updates [][]int
  exampleQueue := getQueue()

  for _, pageQue := range exampleQueue {
    var iq []int
    for i := range pageQue {
      num, _ := strconv.Atoi(pageQue[i])
      iq = append(iq, num)
    }
    updates = append(updates, iq)
  }

  return updates
}

func getRules() [][2]int {
  var rules [][2]int

  file, err := os.Open("./rules.txt")
  if err != nil {
    fmt.Printf("Error opening file: %v\n", err)
    return rules
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scanner.Text()
    parts := strings.Split(line, "|")
    a, _ := strconv.Atoi(parts[0])
    b, _ := strconv.Atoi(parts[1])
    rules = append(rules, [2]int{a, b})
  }

  return rules
}

func getQueue() [][]string {
  var qList [][]string

  file, err := os.Open("queue.txt")
  if err != nil {
    fmt.Printf("Error opening file: %v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    qList = append(qList, strings.Split(scanner.Text(), ","))
  }
  return qList
}
