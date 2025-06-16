package main
import "fmt"
func changeSlice(p []int) []int {
	// Modify the slice by appending a new value
	p[0]=10
	p = append(p, 11)
	 return p
}
func main() {
	// Create a slice of integers
	x:= []int{1, 2, 3,4,5}
	x=append(x, 6)
	x=append(x, 7)

	a:=x[4:]
	y := changeSlice(a)
	fmt.Println("Original slice:", x) // Output: Original slice: [1 2 3 4 5 6 7]
    fmt.Println("Modified slice:", y) //Modified slice: [10 6 7 11]

	fmt.Println(x[0:8]) //[1 2 3 4 10 6 7 11]
}