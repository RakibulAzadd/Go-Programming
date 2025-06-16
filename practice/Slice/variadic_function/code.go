package main
import "fmt"

func print(nums ...int){
	fmt.Println(nums)
	fmt.Println("Length of nums:",len(nums))
	fmt.Println("capacity of nums:", cap(nums))
}

func main() {
	// Create a slice of integers
	print(1, 2, 3, 4, 5) // Passing variadic arguments

}