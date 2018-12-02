package main
import (
  "fmt"
  "github.com/AnjuAshok/slice-in-ascending/sort"
)

func main() {
  var y []int
  var l int
  fmt.Println("Plerase enter the size of array: ")
  fmt.Scanln(&l)
  fmt.Println("Enter the values: ")
  for i:=0; i < l; i++ {
      var temp int
      fmt.Scanln(&temp)
      y = append(y, temp)
  }
	new := sort.Sort(y)
	fmt.Println("The sorted slice is :", new)
}
