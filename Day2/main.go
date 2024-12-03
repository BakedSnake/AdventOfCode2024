//The unusual data (your puzzle input) consists of many reports, one report per line. Each report is a list of numbers called levels that are separated by spaces. For example:
//7 6 4 2 1
//1 2 7 8 9
//9 7 6 2 1
//1 3 2 4 5
//8 6 4 4 1
//1 3 6 7 9
//This example data contains six reports each containing five levels.

//The engineers are trying to figure out which reports are safe. The Red-Nosed reactor safety systems can only tolerate levels that are either gradually increasing or gradually decreasing. So, a report only counts as safe if both of the following are true:

//    The levels are either all increasing or all decreasing.
//    Any two adjacent levels differ by at least one and at most three.

//In the example above, the reports can be found safe or unsafe by checking those rules:

//    7 6 4 2 1: Safe because the levels are all decreasing by 1 or 2.
//    1 2 7 8 9: Unsafe because 2 7 is an increase of 5.
//    9 7 6 2 1: Unsafe because 6 2 is a decrease of 4.
//    1 3 2 4 5: Unsafe because 1 3 is increasing but 3 2 is decreasing.
//    8 6 4 4 1: Unsafe because 4 4 is neither an increase or a decrease.
//    1 3 6 7 9: Safe because the levels are all increasing by 1, 2, or 3.

//So, in this example, 2 reports are safe.

//Analyze the unusual data from the engineers. How many reports are safe?

//--- Part Two ---

//The engineers are surprised by the low number of safe reports until they realize they forgot to tell you about the Problem Dampener.

//The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!

//Now, the same rules apply as before, except if removing a single level from an unsafe report would make it safe, the report instead counts as safe.

//More of the above example's reports are now safe:

//    7 6 4 2 1: Safe without removing any level.
//    1 2 7 8 9: Unsafe regardless of which level is removed.
//    9 7 6 2 1: Unsafe regardless of which level is removed.
//    1 3 2 4 5: Safe by removing the second level, 3.
//    8 6 4 4 1: Safe by removing the third level, 4.
//    1 3 6 7 9: Safe without removing any level.

//Thanks to the Problem Dampener, 4 reports are actually safe!

//Update your analysis by handling situations where the Problem Dampener can remove a single level from unsafe reports. How many reports are now safe?


package main

import (
  "fmt"
  "strings"
  "strconv"
)

var Report = []string{
  "7 6 4 2 1",
  "1 2 7 8 9",
  "9 7 6 2 1",
  "1 3 2 4 5",
  "8 6 4 4 1",
  "1 3 6 7 9",
  "10 30 66 70 90",
  "10 12 15 17 90",
  "46 47 50 54 57 60 63 67",
}

func main() {
  count := 0
  for i := range Reports {
    originalReport := getSingleReport(Reports[i])
    if checkSafe(originalReport) {
      fmt.Printf("%v: Is safe!!!\n", originalReport)
      count++
    } else {
      // second part
      if checkSafeRemove(originalReport) {
        fmt.Printf("%v: Is safe after removing!!!\n", originalReport)
        count++
      } else {
        fmt.Printf("%v: Not safe...\n", originalReport)
      }
    }
  }
  fmt.Println("----------")
  fmt.Println("----------")
  fmt.Println("----------")
  fmt.Println("----------")
  fmt.Println(count)
}

//second part
func checkSafeRemove(report []int) bool {
  for i := 0; i < len(report); i++ {
    newReport := append([]int{}, report[:i]...)
    newReport = append(newReport, report[i+1:]...)
    if checkSafe(newReport) {
      return true
    }
  }
  return false
}

func checkSafe(r []int) bool {
  isIncreasing := true
  isDecreasing := true
  for i := 0; i< len(r)-1; i++ {
    diff := r[i+1] - r[i]
    if diff >= 0 {
      isDecreasing = false
    }
    if diff <= 0 {
      isIncreasing = false
    }
  }
  if !isIncreasing && !isDecreasing {
    return false
  }
  for i := 0; i < len(r)-1; i++ {
    diff := r[i+1] - r[i]
    if diff < -3 || diff > 3 || diff == 0 {
      return false
    }
  }
  return true
}

func getSingleReport(r string) []int {
  var reportDigits []int
  digits := strings.Split(r, " ")
  for i := range digits {
    idig, _ := strconv.Atoi(digits[i])
    reportDigits = append(reportDigits, idig)
  }
  return reportDigits
}
