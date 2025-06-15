package main
import "fmt"
var love = [3]string{"I","Love","you"} // Global variable

func main(){
   var ar [2]int 
   ar[0]=2
   ar[1]=3
   fmt.Println("Array:", ar)
   
   arr := [2]int{1, 2}
   fmt.Println("Array:", arr)

   fmt.Println(love)


}
