package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{2, 4, 6, 4, 2, 3, 4, 1, 5, 6}  
	A, B, C := 2, 4, 6                          
   
	// Time limit = 10;
	 
	mp := make(map[int]int)

	 
	for _, people := range arr {
		if people == 1 {
			people = 2
		} else if people == 3 {
			people = 4
		} else if people == 5 {
			people = 6
		}
		mp[people]++
	}

	 
	mn := math.MaxInt32
	for _, count := range mp {
		if count < mn {
			mn = count
		}
	}

	 
	mp[A] -= mn
	mp[B] -= mn
	mp[C] -= mn

	unsetted := 0

	 
	for {  

		if mp[A] == 0 && mp[B] == 0 && mp[C] == 0 {
			break
		}

		if mp[C] > 0 {
			mp[C]--
		} else if mp[B] > 0 {
			mp[B]--
		} else if mp[A] > 0 {
			mp[A]--
		} else {
			unsetted++
		}

		if mp[B] > 0 {
			mp[B]--
		} else if mp[A] > 0 {
			mp[A]--
		} else {
			unsetted++
		}

		if mp[A] > 0 {
			mp[A]--
		} else {
			unsetted++
		}
	}

	fmt.Println("Unseated Table:", unsetted)
}
