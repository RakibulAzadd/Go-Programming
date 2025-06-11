package main 
import "fmt"
func main(){
	var a,b int = 10,20

	if a<b {
		fmt.Printf("%d is less than %d\n", b,a)
	}  else {
		fmt.Printf("%d is greater than %d\n", b,a);
	}
	fmt.Println("End of program");
}