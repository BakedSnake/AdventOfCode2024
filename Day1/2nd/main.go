package main

import (
  "fmt"
)

func main() {
  fmt.Println(sumApps(countCheck(listA, listB)))
}

// Similarity score is the sum of all appearances of listA values in listB
func sumApps(arr []int) int {
  var simScore int
  for i := range arr {
    simScore += arr[i]
  }
  return simScore
}

func countCheck(arr1 []int, arr2 []int) []int {
  var appList []int
  for i := range arr1 {
    count := 0
    for j := range arr2 {
      if arr1[i] == arr2[j] {
        count++
      }
    }
    appList = append(appList, arr1[i] * count)
  }
  return appList
}
