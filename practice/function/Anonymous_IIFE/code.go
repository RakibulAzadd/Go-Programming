package main

import "fmt"

func main() {
	// Anonymous IIFE (Immediately Invoked Function Expression)
	func() {
		fmt.Println("This is an anonymous IIFE function")
	}()

	// Another example of an IIFE with parameters
	func(x int, y int) {
		sum := x + y
		fmt.Println("Sum:", sum)
	}(5, 10) // Immediately passing arguments
}

func init(){
	fmt.Println("This is the init function")
	// This function runs before main
	// It can be used to set up initial state or configurations	
}