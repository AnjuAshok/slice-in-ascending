package sort

func Sort(x []int) []int{
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
