package main

import "fmt"

func main() {
	

	//add(2,3) // Calling the function add and error for not defining it yet

	add:=func(x int, y int){
		c:=x+y
		fmt.Println("The sum is:", c)
	}

	add(5,6)
	 
}