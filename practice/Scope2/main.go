package main
import "fmt"
func l_add(x int, y int)int{
	sum:=x+y
	return sum
}
func main(){
	p:=30
	q:=40
	sum:=l_add(p,q)
	fmt.Println(sum)
	dd:=add(p,q)
	fmt.Println(dd)
}