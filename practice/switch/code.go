package main
import "fmt"
func main() {
	var a,b int = 10,20
	switch{
	case a==b:
		fmt.Printf("%d is equal to %d\n", a, b)
	case a<b:
		fmt.Printf("%d is less than %d\n", a, b)
	default:
		fmt.Printf("%d is greater than %d\n", a, b)	
	}

	switch a {
	case  10:
		fmt.Println("a is 10")
	case 20:
		fmt.Println("a is 20")
	
	}
}