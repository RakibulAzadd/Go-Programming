package main
import "fmt"
func main() {
	var i int
	for i = 1; i <= 5; i++ {
		fmt.Printf("i is %d\n", i)
	}
	
	fmt.Println("End of for loop")
	
	j := 1
	for j <= 5 {
		fmt.Printf("j is %d\n", j)
		j++
	}
	
	fmt.Println("End of while loop")
	
	k := 1
	for {
		if k > 5 {
			break
		}
		fmt.Printf("k is %d\n", k)
		k++
	}
	
	fmt.Println("End of infinite loop")
}
