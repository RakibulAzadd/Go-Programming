package main

import "fmt"

func processOperation(a int , b int, op func(p int, q int)){
	op(a,b)
}
func add(x int, y int){
	z:=x+y
	fmt.Println(z)
}
func main() {
	processOperation(2,5,add)
}

//high order function
// callback function
//first class citizen
