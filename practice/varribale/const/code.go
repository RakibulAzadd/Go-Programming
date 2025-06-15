package main
import "fmt"
var g =20
const cs= 40
func call(a int) {
	const c= 1
	fmt.Println("a:", a)	
}
func main() {
	var a int =10
	call(a)
	fmt.Println("a:", a)
	fmt.Println("g:", g)
	fmt.Println("cs:", cs)
	fmt.Println("c:", 1) // This will print the constant value directly
}