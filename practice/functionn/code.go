package main
import "fmt"
func add(num1 int, num2 int){
	sum:=num1+num2
	fmt.Println(sum)
}
func return_add(num1 int, num2 int) int{
	sum:= num1 + num2

	return sum
}
func return_multiple_value(num1 int, num2 int) (int,int){
	sum:= num1 + num2
    mul := num1*num2
	return sum,mul
}
func say_hello(){
	fmt.Println("Welcome to golang course ")
}
func hi_rakib(name string){
	fmt.Println("Welcome to you, ", name)
}
func main() {	
	a:= 10
	b:= 20

	add(a,b)
    R_sum := return_add(a, b)
	fmt.Println(R_sum)

	p,q := return_multiple_value(a,b)

	fmt.Println(p,q)
	say_hello()
	hi_rakib("Rakib")
	

}