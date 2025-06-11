package main
import "fmt"

func main(){
	var a,b int
	a=30
	b=90
	
	add(a,b)
	fmt.Println("code finish")
}
func printfun(sum int ){
	fmt.Println(sum)
}
func add(x int ,y int){
    sum:= x+y
	printfun(sum)
}