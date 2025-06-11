package main
import "fmt"
var a=20
var b=30

func add(x int, y int){
	sum:= x+y
	fmt.Println(sum)
}

func main(){
	p:=40
	q := 50

	add(p,q)
	add(a,p)
	add(b,q)

	// add(x,a) scope error 	fmt.Print("Hello") // Output: Hello
	fmt.Print("Hello") // Output: Hello
	fmt.Print("World") // Output: HelloWorld
}