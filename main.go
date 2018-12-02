package main
import "fmt"

func sort(x []int) []int{
  for i := 0; i < len(x); i++ {
    for j := i+1; j<len(x); j++ {
      if x[i] > x[j]{
        a := x[i]
        x[i] = x[j]
        x[j] = a
      }
    }
  }
	return x
}

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
	new := sort(y)
	fmt.Println("The sorted slice is :", new)
}
