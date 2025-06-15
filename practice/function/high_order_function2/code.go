package main

import "fmt"

 func call() func(x int, y int){
	return add
 }

func add(x int, y int) {	
	sum := x + y
	fmt.Println(sum)
}

func main() {
 sum:= call()
 sum(4,3)
 fmt.Println("finish")
}
