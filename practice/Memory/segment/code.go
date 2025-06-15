package main
import "fmt"
var a int = 10
const p=30
func call() {
	add := func(x int, y int) {
		z := x + y
		fmt.Println("Sum:", z)
	}
	add(5,6)
	add(p,a)
}
func main() {
	a=40
	call()
	fmt.Println("a:", a)

}
func init() {
	fmt.Println("init function")
}	