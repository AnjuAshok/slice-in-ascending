package main
import "fmt"

func sort (x []int) []int{

        for i:=0; i < l; i++{
                for j:=i+1 ;j<l; j++{
                        if x[i] >x[j]{
                        a := x[i]
                        x[i]=x[j]
                        x[j] = a
                        }
                }
	return x[i]
}
func main() {
	var y := make([]int,5)
	l := len(y)
         for i:=0; i < l ;i++{
                fmt.Scanln(&y[i])
        }
	new := sort(y)
	fmt.Println("The sorted slice is :", new)




}
