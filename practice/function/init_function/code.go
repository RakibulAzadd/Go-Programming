package main

import "fmt"

var a = 30

func main(){
	fmt.Println(a)
}
func init(){
	fmt.Println(a)
	a = 40
}
