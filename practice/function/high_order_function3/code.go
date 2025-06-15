package main

import "fmt"

func processOperation(a int , b int, op func(p int, q int)) func(p int,q int){
	op(a,b)

	return add
}
func add(x int, y int){
	z:=x+y
	fmt.Println(z)
}
func main() {
	sum:=processOperation(2,5,add)

	sum(1,2)
	fmt.Println("code ses")
}

//high order function
// callback function
//first class citizen
