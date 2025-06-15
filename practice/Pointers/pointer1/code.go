package main
import "fmt"

type User struct {
	Name string	
	Age  int
}

func print(arr *[3]int) {
	fmt.Println("Array:", arr)
}
func main() {
	var a int = 10
	 
	p := &a // p is a pointer to a
 
	fmt.Println("Value pointed by p:", *p) // Output: Value pointed by p: 10
	 

	arr := [3]int{1, 2, 3}

	print(&arr)

	usr := User{
		Name: "Habib",	
		Age:  23,
	}

	pp := &usr // pp is a pointer to usr

	fmt.Println(*pp)

	fmt.Println(pp.Name) // Accessing field through pointer

	 
}