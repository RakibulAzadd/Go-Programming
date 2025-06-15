package main
import "fmt"
type User struct {
	Name string
	Age  int
}
func printUserDetails(usr User) {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}	
func (usr User) printDetails() {
	fmt.Println("Name:", usr.Name)
	fmt.Println("Age:", usr.Age)
}	
func (usr User) call(a int){
	fmt.Println("Name:", usr.Name)
	fmt.Println(a)
}
func main() {
	var user1 User
	user1 = User{
		Name: "Habib",
		Age:  23,
	}
	user2 := User{
		Name: "Rakib",
		Age:  21,
	}

	printUserDetails(user1)
	printUserDetails(user2)

	user1.printDetails()
	user2.printDetails()

	user1.call(10)
}