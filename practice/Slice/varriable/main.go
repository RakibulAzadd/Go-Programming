package main
import (
	"fmt"
)
func main() {
	// Create a slice of integers
	var x []int	
	 x = append(x, 1)
	s := make([]int, 4, 5)
	
	nums:= []int{1,2,4,5,6}
	fmt.Println("nums:", nums)
	fmt.Println("x:", x)
    fmt.Println("s:", s)	
}