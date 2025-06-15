package main
import "fmt"

type User struct{
	Name string
	Age int
}
func main(){
	var user1 User
	user1=User{
		Name: "Habib",
		Age: 23,
	}
	user2 := User{
		Name: "Rakib",
		Age: 21,
	}

	fmt.Println("Name : ", user1.Name)
	fmt.Println("Age : ", user1.Age)

	fmt.Println("Name : ", user2.Name)
	fmt.Print("Age : ", user2.Age)

}
func init() {
	fmt.Println("This is the init function")
}