package main
import ("fmt"
 "example.com/mathlib")
var (
	a=10
	b=20
)
func l_add(x int, y int)int{
	sum:=x+y
	return sum
}
func main(){
	p:=30
	q:=40
	sum:=l_add(a,b)
	fmt.Println(sum)
	dd:=mathlib.Add(p,q)
	fmt.Println(dd)
}

 