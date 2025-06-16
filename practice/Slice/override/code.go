package main

import (
	"fmt"
)
func main() {
	// Create a slice of integers
	var x []int 

	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	
	y := x

	x = append(x, 4)
	fmt.Println("x:", x)
	
	y = append(y, 5)
    fmt.Println("y:", y)
}