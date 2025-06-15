package main
import "fmt"
var a int = 10
func main() {
	 
	var b int = 20

	if a < b {
		 a=20
	}  

	if a == 10 {
		fmt.Println("a is 10")
	} else if a == 20 {
		fmt.Println("a is 20")
	} else {
		fmt.Println("a is neither 10 nor 20")
	}
}