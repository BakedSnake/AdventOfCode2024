package main

import (
  "fmt"
)

func main() {
  arr1 := quickSortStart(listA)
  arr2 := quickSortStart(listB)
  diffs := getDifference(arr1, arr2)
  fmt.Println(sumDiff(diffs))
}

func sumDiff(d []int) int {
  sum := 0
  for i := range(d)  {
    sum += d[i]
  }
  return sum
}

func getDifference(arr1 []int, arr2 []int) []int {
  var diffs []int
  for i := range arr1 {
    diff := arr1[i] - arr2[i]
    if diff < 0 {
      diff = -diff
    }
    diffs = append(diffs, diff)
  }
  return diffs
}

func quickSortStart(arr []int) []int {
  return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low, high int) []int {
  if low < high {
    var p int
    arr, p = partition(arr, low, high)
    arr = quickSort(arr, low, p-1)
    arr = quickSort(arr, p+1, high)
  }
  return arr
}

func partition(arr []int, low, high int) ([]int, int) {
  pivot := arr[high]
  i := low
  for j := low; j < high; j++ {
    if arr[j] < pivot {
      arr[i], arr[j] = arr[j], arr[i]
      i++
    }
  }
  arr[i], arr[high] = arr[high], arr[i]
  return arr, i
}
