package main

import (
  "os"
  "fmt"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  ruleParser()
}

func ruleParser() [][]int {
  var rules [][]int
  var first, after int
  exampleRules := getRules()

  fmt.Printf("[RULES]:\n")
  for _, rule := range exampleRules {
    var r []int
    pages := strings.Split(string(rule), "|")
    first, _ = strconv.Atoi(pages[0])
    after, _ = strconv.Atoi(pages[1])
    r = append(r, first)
    r = append(r, after)
    rules = append(rules, r)
    fmt.Printf("%v\n", r)
  }
  return rules
}

func getRules() [][]rune {
  var qList [][]rune

  file, err := os.Open("./example-rules.txt")
  if err != nil {
    fmt.Printf("%v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    qList = append(qList, []rune(scanner.Text()))
  }
  return qList
}

func getQueue() [][]rune {
  var qList [][]rune

  file, err := os.Open("./example-queue.txt")
  if err != nil {
    fmt.Printf("%v", err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    qList = append(qList, []rune(scanner.Text()))
  }
  return qList
}
